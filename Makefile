all: build

prepare:
	@if which node; then\
		echo "Node.JS installed...";\
	else\
		echo "Node.JS is not installed!";\
		exit 1;\
	fi
	@if which npm; then\
		echo "NPM installed...";\
	else\
		echo "NPM is not installed!";\
		exit 1;\
	fi
	@echo "Installing Gulp CLI..."
	@sudo npm i -g gulp-cli
	@echo "Installing go-bindata..."
	@go get github.com/jteeuwen/go-bindata/...
	@go get github.com/elazarl/go-bindata-assetfs/...
	@echo "Installing Go dependencies..."
	@go get -t ./...
	@echo "Installing Node.JS dependencies..."
	@npm i

prepare-ci:
	@npm i -g gulp-cli 
	@go get github.com/jteeuwen/go-bindata/...
	@go get github.com/elazarl/go-bindata-assetfs/...
	@go get -t ./...
	@npm i

web:
	@gulp
	@go-bindata-assetfs dist

build: web
	@go build

run: web
	@go run settings.go bindata.go validate.go response.go fetch.go check.go up.go

docker: web
	@docker build -t is-up .
	
test:
	@go test

clean:
	@rm -rf dist up
