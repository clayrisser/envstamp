.PHONY: all
all: clean build

.PHONY: start
start:
	@go run envstamp.go template.txt

.PHONY: build
build: envstamp
	@echo ::: BUILD :::

.PHONY: clean
clean:
	-@rm -rf envstamp template.txt &>/dev/null || true
	@echo ::: CLEAN :::

.PHONY: install
install: build
	@mv envstamp /bin
	@chmod +x /bin/envstamp
	@echo ::: INSTALL :::

envstamp:
	@go build envstamp.go
