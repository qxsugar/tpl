#! /bin/bash

dest="${1:-app}"
PROJECT_DIR=.tmp/
rm -rf $PROJECT_DIR
mkdir -p $PROJECT_DIR
git clone git@github.com:qxsugar/tpl.git $PROJECT_DIR
go-archetype transform --transformations=${PROJECT_DIR}transformation-api.yml --source=${PROJECT_DIR}src/api --destination=${dest}
rm -rf $PROJECT_DIR
echo "done"