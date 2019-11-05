build_binary () {
    if [ -z "$1" ]
    then
        echo "You must specify a platform.\n./build.sh [darwin | linux | windows]"
    else
        env GOOS=$1 GOARCH=amd64 go build -ldflags="-s -w" -o keep *.go
        echo "Built binaries for $1".

        if [ "$2" = "zip" ]
        then
            zip -r keep-$1.zip keep
            echo "Zipped binary as keep-$1.zip."
        else
            echo "Built binaries. Did not ZIP."
        fi
    fi
}

build_binary $1 $2