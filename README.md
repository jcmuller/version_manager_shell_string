# version_manager_shell_string : Shell string with version manager strings

## Overview
[![GoDoc](https://godoc.org/github.com/jcmuller/version_manager_shell_string?status.svg)](https://godoc.org/github.com/jcmuller/version_manager_shell_string)
[![Code Climate](https://codeclimate.com/github/jcmuller/version_manager_shell_string/badges/gpa.svg)](https://codeclimate.com/github/jcmuller/version_manager_shell_string)
[![Go Report Card](https://goreportcard.com/badge/github.com/jcmuller/version_manager_shell_string)](https://goreportcard.com/report/github.com/jcmuller/version_manager_shell_string)
[![Sourcegraph](https://sourcegraph.com/github.com/jcmuller/version_manager_shell_string/-/badge.svg)](https://sourcegraph.com/github.com/jcmuller/version_manager_shell_string?badge)
[![Build Status](https://travis-ci.org/jcmuller/version_manager_shell_string.svg?branch=master)](https://travis-ci.org/jcmuller/version_manager_shell_string)

Configurable single binary to gather all your version managers active versions.

## Install

```
go get github.com/jcmuller/version_manager_shell_string
```

## Example

Set up `~/.config/version_manager_shell_string/config.yml`:
```yaml
- command: rbenv
  args:
  - version-name
  identifier: R
  file: .ruby-version
- command: goenv
  args:
  - version-name
  identifier: G
  file: .go-version
```

```bash
$ version_manager_shell_string
R:2.4.3*|N:v8.6.0*|G:system
```

That means:
- rbenv: 2.4.3
- nvm: 8.6.0
- goenv: system
```

## License

MIT.
