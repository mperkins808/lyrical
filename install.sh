#!/bin/bash

GITHUB_REPO="mperkins808/lyrical"
BINARY_NAME="lyrical"
TAG="0.0.3"
BINARY_URL="https://github.com/$GITHUB_REPO/releases/download/$TAG/$BINARY_NAME"

TEMP_DIR=$(mktemp -d)

curl -L $BINARY_URL -o $TEMP_DIR/$BINARY_NAME

chmod +x $TEMP_DIR/$BINARY_NAME

cp $TEMP_DIR/$BINARY_NAME /usr/local/bin/$BINARY_NAME

rm -rf $TEMP_DIR

echo "$BINARY_NAME installed to /usr/local/bin"