APPNAME := "xbar_github_status.5m.cgo"

.PHONY: build
build:
	go build -o bin/$(APPNAME) .
	chmod +x bin/$(APPNAME)

.PHONY: clean
clean:
	rm -rf bin