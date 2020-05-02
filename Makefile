# The root package
PKGROOT=github.com/Ladicle/git-prompt

# OUTDIR is directory where artifacts
OUTDIR=_output

build:
	CGO_ENABLED=0 go build -o $(OUTDIR)/git-prompt

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -o $(OUTDIR)/git-prompt_linux-amd64/git-prompt

build-darwin:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 \
	go build -o $(OUTDIR)/git-prompt_darwin-amd64/git-prompt

build-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
	go build -o $(OUTDIR)/git-prompt_windows-amd64/git-prompt

install:
	CGO_ENABLED=0 go install

.PHONY: clean
clean:
	-rm -r $(OUTDIR)
