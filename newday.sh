#!/usr/bin/env sh
if [ $# -ne 2 ]; then
    echo "$0 <year> <day>"
    exit 1
fi

dir="$1/$2"

mkdir -p $dir
touch "$dir/$2.go"
touch "$dir/example"
touch "$dir/input"