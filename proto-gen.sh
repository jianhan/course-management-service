#!/bin/bash
echo start to generate protobuf
for file in `find | grep -E "\.proto$"`; do
echo generating "$file"
protoc --go_out=plugins=micro,grpc:. $file
done
