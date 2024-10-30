#!/bin/bash
RUN_NAME=hertz_service
mkdir -p output/bin
cp -r conf/ output/
cp -r docs/ output/

go build -o output/bin/${RUN_NAME}