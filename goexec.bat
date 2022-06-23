@ECHO OFF

set GOBIN=./bin
set APP_NAME=task2
set arg1=%1

GOTO %1

:build
    @go build -o %GOBIN%/%APP_NAME%.exe ./cmd/%APP_NAME%/main.go
    @echo Done building.
	@echo Run %GOBIN%/%APP_NAME% to launch %APP_NAME%.
exit

:run
    @go run ./cmd/%APP_NAME%/main.go
	@echo Done running
exit

:test
    @go test ./... -v -coverprofile=coverage.out && go tool cover -func=coverage.out
exit
