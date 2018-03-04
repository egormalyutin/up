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

build:
	@gulp
	@go-bindata-assetfs dist
	@go build

run: build
	@./up

clean:
	@rm -rf dist up
