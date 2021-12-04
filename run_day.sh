mkdir -p -- "$1"
cd $1
echo "Running task: $1"
if [ "$#" -eq 2 ]; then
    if [ "$2" == "init" ]; then
        go mod init $1
    fi
fi
go build
time cat input | ./$1