# The root package
PKGROOT=github.com/Ladicle/git-prompt

# OUTDIR is directory where artifacts
OUTDIR=_output

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(OUTDIR)/git-prompt

.PHONY: install
install:
	CGO_ENABLED=0 go install

.PHONY: clean
clean:
	-rm -r $(OUTDIR)
