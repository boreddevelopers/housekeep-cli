env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o keep *.go
zip -r keep-darwin.zip keep
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o keep *.go
zip -r keep-linux-amd64.zip keep