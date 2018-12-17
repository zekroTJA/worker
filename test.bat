@ECHO OFF

IF "%GOROOT%" == "" (
    ECHO GOROOT envorment variable is not set!
)

go test -v -cover -timeout 30s