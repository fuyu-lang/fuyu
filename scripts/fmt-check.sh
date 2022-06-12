#!/usr/bin/env sh

# Check that all files are properly formatted.

files="$(gofmt -l .)"
if [ ! -z "$files" ]; then
  echo "$files"
  exit 1
fi
