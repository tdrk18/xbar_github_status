APPNAME := "xbar_github_status.5m.cgo"

.PHONY: build
build:
	go build -o bin/$(APPNAME) .
	chmod +x bin/$(APPNAME)

.PHONY: xbuild
xbuild:
	GOOS=darwin; GOARCH=amd64; CGO_ENABLED=0; go build -o bin/$$GOOS/$(APPNAME) -v .; chmod +x bin/$$GOOS/$(APPNAME)
	GOOS=linux; GOARCH=amd64; CGO_ENABLED=0; go build -o bin/$$GOOS/$(APPNAME) -v .; chmod +x bin/$$GOOS/$(APPNAME)
	GOOS=windows; GOARCH=amd64; CGO_ENABLED=0; go build -o bin/$$GOOS/$(APPNAME) -v .; chmod +x bin/$$GOOS/$(APPNAME)

.PHONY: archive
archive:
	@if [ -e bin ]; then \
		zip -r artifacts.zip bin; \
	fi

.PHONY: clean
clean:
	rm -rf bin artifacts.zip