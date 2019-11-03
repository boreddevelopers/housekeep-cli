# env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o keep *.go
env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o keep *.go
# env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o keep *.go