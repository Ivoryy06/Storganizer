BINARY_NAME=storganizer
OUT_DIR=bin

build:
	mkdir -p $(OUT_DIR)
	# Linux build
	GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/$(BINARY_NAME) main.go
	# Windows build
	GOOS=windows GOARCH=amd64 go build -o $(OUT_DIR)/$(BINARY_NAME).exe main.go

clean:
	rm -rf $(OUT_DIR)
	rm storganizer
