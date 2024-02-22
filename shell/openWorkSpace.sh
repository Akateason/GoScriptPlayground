#!/bin/sh
for file in *; do
    if [[ $file == *.xcworkspace ]]; then
        open $file
    fi
done
