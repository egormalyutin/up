# up

## Compilation
Install Go and add GOPATH to your PATH.

### Compile bindata
Install bindata tool:
```bash 
go get github.com/jteeuwen/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
```

Install Node.JS, NPM, then install Gulp and module dependencies:
```bash
(sudo) npm i -g gulp-cli
npm i
```

Compile bindata:
```bash
gulp
go-bindata-assetfs dist
```

### Compile binary
```bash
go build
```

Or run:
```bash
go run bindata.go settings.go validate.go check.go up.go
```
