#!/bin/bash
for file in $(ls -F | grep "/$")
do
    cd $file
    go test -v
    cd ..
done