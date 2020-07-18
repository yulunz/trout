#!/usr/bin/env bash

workingDir="$(dirname $0)"
mkdir -p $workingDir/dist/static

cp cat_fish.jpg $workingDir/dist/static
cp favicon.ico $workingDir/dist/static
cp trout.html $workingDir/dist/index.html
