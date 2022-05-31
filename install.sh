#! /bin/bash

dest="${1:-app}"
PROJECT_DIR=.tmp/
rm -rf $PROJECT_DIR
mkdir -p $PROJECT_DIR
git clone git@github.com:qxsugar/tpl.git $PROJECT_DIR
make generate-api desc=$dest -f $PROJECT_DIR/Makefile
rm -rf $PROJECT_DIR
echo "done"