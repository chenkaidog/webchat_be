#!/bin/bash
RUN_NAME=web_chat
mkdir -p output/bin
cp -r conf/ output/

go build -o output/bin/${RUN_NAME}