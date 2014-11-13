#!/bin/sh
rm -rf docs
for package in $(go list ./...)
do
    DIR=$(echo $package | sed 's%github.com/gtarcea/1DevDayTalk2014%.%')
    mkdir -p docs/$DIR
    godoc $package > docs/$DIR/package.txt
done
