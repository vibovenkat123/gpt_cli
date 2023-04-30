BIN=./bin
FILE=gptcli
SRC=cmd/cli/main.go
BUILD_SH=./scripts/build.sh

default: local
local: clean fmt ensure_deps lint fix_ordering build
distrobute: local buildall

lint:
	$(info ******** Linting the app ********)
	golangci-lint run ./...
fmt:
	$(info ******** Formatting the app ********)
	gofmt -w -s -l .
fix_ordering:
	$(info ******** Fixing the struct ordering ********)
	fieldalignment -fix ./...
clean:
	$(info ********** CLEANING BIN **********)
	rm -rf $(BIN)
build:
	$(info ********** BUILDING CLI **********)
	go build -o $(BIN)/$(FILE) $(SRC)
ensure_deps:
	go mod tidy
buildall:
	$(info ********** BUILDING ALL **********)
	GOOS=darwin GOARCH=arm64 $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="darwin" GOARCH="amd64" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="linux" GOARCH="amd64" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="linux" GOARCH="386" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="linux" GOARCH="arm" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="linux" GOARCH="arm64" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="linux" GOARCH="ppc64" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="linux" GOARCH="ppc64le" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="linux" GOARCH="mips" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="linux" GOARCH="mipsle" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="linux" GOARCH="mips64" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="linux" GOARCH="mips64le" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="windows" GOARCH="386" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="windows" GOARCH="amd64" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
	GOOS="windows" GOARCH="arm64" $(BUILD_SH) $(BIN) $(FILE) $(SRC)
