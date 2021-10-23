#!/usr/bin/env bash

go build -o example
chmod +x example
mkdir -p ~/.gauge/plugins/example/0.1.0
mv example ~/.gauge/plugins/example/0.1.0
cp plugin.json ~/.gauge/plugins/example/0.1.0
