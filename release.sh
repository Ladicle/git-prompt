#!/bin/bash -ex

# cleanup output directory
make clean

# build binary
make build-linux
make build-darwin
make build-windows

# compose binaries
cd _output/
for dir in $(find . -type d -name "git-prompt_*"); do
    tar -zcvf $dir.tar.gz $dir
done

# check sha256
sha256sum *.tar.gz
