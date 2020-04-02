all: build

cmd/version_manager_shell_string/version_manager_shell_string: cmd/version_manager_shell_string/main.go internal/checker/checker.go internal/config/config.go internal/versions/versions.go
	cd cmd/version_manager_shell_string; CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"'

.PHONY: build
build: cmd/version_manager_shell_string/version_manager_shell_string

.PHONY: clean
clean:
	rm cmd/version_manager_shell_string/version_manager_shell_string

.PHONY: update-dependencies
update-dependencies:
	go get -u -v ./...
	go mod tidy
	go mod verify
	go mod vendor
