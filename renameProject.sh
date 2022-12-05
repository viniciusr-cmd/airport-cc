#!/usr/bin/env bash

# CAUTION: this script will replace every occurrence of the word
# `airport-cc` in the project folder with whatever argument
# you pass. Be very careful.

if [ $# -lt 1 ] ; then
  printf "Usage:\n$ ./renameProject.sh <my-project-name>\n"
  exit
fi

grep -rl airport-cc . --exclude-dir={.git,node_modules} | xargs sed -i s/airport-cc/$1/g
