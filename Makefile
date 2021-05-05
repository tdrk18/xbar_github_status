APPNAME := "xbar_github_status.5m.cgo"

.PHONY: build
build:
	go build -o bin/$(APPNAME) .
	chmod +x bin/$(APPNAME)

.PHONY: xbuild
xbuild:
	export GOOS=darwin; export GOARCH=amd64; export CGO_ENABLED=0; go build -o bin/$$GOOS/$(APPNAME) .; chmod +x bin/$$GOOS/$(APPNAME)
	export GOOS=linux; export GOARCH=amd64; export CGO_ENABLED=0; go build -o bin/$$GOOS/$(APPNAME) .; chmod +x bin/$$GOOS/$(APPNAME)
	export GOOS=windows; export GOARCH=amd64; export CGO_ENABLED=0; go build -o bin/$$GOOS/$(APPNAME) .; chmod +x bin/$$GOOS/$(APPNAME)

.PHONY: archive
archive:
	@if [ -e bin ]; then \
		zip -r artifacts.zip bin; \
	fi

.PHONY: clean
clean:
	rm -rf bin artifacts.zip