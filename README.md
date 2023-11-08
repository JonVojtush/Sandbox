go.mod
    go mod vendor   or   go get {src}
    go mod tidy
    go clean -modcache

go.work
    https://github.com/golang/tools/blob/master/gopls/doc/workspace.md

go build vs go install:
    go build just compiles the executable file and moves it to the destination. go install does a little bit more. It moves the executable file to $GOPATH/bin and caches all non-main packages which are imported to $GOPATH/pkg. The cache will be used during the next compilation provided the source did not change yet.
    
    ├── bin
    │   └── hello  # by go install
    └── src 
        └── hello
            ├── hello  # by go build
            └── hello.go