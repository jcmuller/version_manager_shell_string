package nvmchecker

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// NvmChecker
type NvmChecker struct {
	basePath   string
	cmd        *exec.Cmd
	version    string
	defined    bool
	identifier string
}

var (
	file     = ".nvmrc"
	command  = "bash"
	argument = "/tmp/nvm_checker.sh"
	script   = `
#!/bin/bash
nvm_ls_current() {
  local NVM_LS_CURRENT_NODE_PATH
  if ! NVM_LS_CURRENT_NODE_PATH="$(command which node 2> /dev/null)"; then
	nvm_echo 'none'
  elif nvm_tree_contains_path "$(nvm_version_dir iojs)" "${NVM_LS_CURRENT_NODE_PATH}"; then
	nvm_add_iojs_prefix "$(iojs --version 2>/dev/null)"
  elif nvm_tree_contains_path "${NVM_DIR}" "${NVM_LS_CURRENT_NODE_PATH}"; then
	local VERSION
	VERSION="$(node --version 2>/dev/null)"
	if [ "${VERSION}" = "v0.6.21-pre" ]; then
	  nvm_echo 'v0.6.21'
	else
	  nvm_echo "${VERSION}"
	fi
  else
	nvm_echo 'system'
  fi
}
nvm_version_dir() {
  local NVM_WHICH_DIR
  NVM_WHICH_DIR="${1-}"
  if [ -z "${NVM_WHICH_DIR}" ] || [ "${NVM_WHICH_DIR}" = "new" ]; then
	nvm_echo "${NVM_DIR}/versions/node"
  elif [ "_${NVM_WHICH_DIR}" = "_iojs" ]; then
	nvm_echo "${NVM_DIR}/versions/io.js"
  elif [ "_${NVM_WHICH_DIR}" = "_old" ]; then
	nvm_echo "${NVM_DIR}"
  else
	nvm_err 'unknown version dir'
	return 3
  fi
}
nvm_tree_contains_path() {
  local tree
  tree="${1-}"
  local node_path
  node_path="${2-}"

  if [ "@${tree}@" = "@@" ] || [ "@${node_path}@" = "@@" ]; then
	nvm_err "both the tree and the node path are required"
	return 2
  fi

  local pathdir
  pathdir=$(dirname "${node_path}")
  while [ "${pathdir}" != "" ] && [ "${pathdir}" != "." ] && [ "${pathdir}" != "/" ] && [ "${pathdir}" != "${tree}" ]; do
	pathdir=$(dirname "${pathdir}")
  done
  [ "${pathdir}" = "${tree}" ]
}
nvm_err() {
  >&2 nvm_echo "$@"
}
nvm_echo() {
  command printf %s\\n "$*" 2>/dev/null
}
nvm_version() {
  local PATTERN
  PATTERN="${1-}"
  local VERSION
  # The default version is the current one
  if [ -z "${PATTERN}" ]; then
	PATTERN='current'
  fi

  if [ "${PATTERN}" = "current" ]; then
	nvm_ls_current
	return $?
  fi

  local NVM_NODE_PREFIX
  NVM_NODE_PREFIX="$(nvm_node_prefix)"
  case "_${PATTERN}" in
	"_${NVM_NODE_PREFIX}" | "_${NVM_NODE_PREFIX}-")
	  PATTERN="stable"
	;;
  esac
  VERSION="$(nvm_ls "${PATTERN}" | command tail -1 | trim)"
  echo "${VERSION}"
}

nvm_version
`
)

// New instance
func New(path string) *NvmChecker {
	return &NvmChecker{
		basePath:   path,
		identifier: "N",
	}
}

// StartCheck
func (l *NvmChecker) StartCheck() {
	_, err := os.Stat(argument)

	if err != nil {
		file, err := os.Create(argument)
		handle(err)
		_, err = file.WriteString(script)
		handle(err)
		err = file.Sync()
		handle(err)
	}

	l.cmd = exec.Command(command, argument)

	reader, err := l.cmd.StdoutPipe()
	handle(err)

	scanner := bufio.NewScanner(reader)

	go func() {
		for scanner.Scan() {
			l.version = scanner.Text()
		}
	}()

	err = l.cmd.Start()
	handle(err)
}

// Wait
func (l *NvmChecker) Wait() {
	err := l.cmd.Wait()
	handle(err)
}

func (l *NvmChecker) setDefined() {
	file = filepath.Join(l.basePath, file)
	_, err := os.Stat(file)
	l.defined = err == nil
}

// GetVersion does that
func (l *NvmChecker) GetVersion() {
	l.setDefined()
	l.Wait()
}

// Output string
func (l *NvmChecker) String() string {
	str := fmt.Sprintf("%s:%s", l.identifier, l.version)

	if l.defined {
		str = strings.Join([]string{str, "*"}, "")
	}

	return str
}

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
