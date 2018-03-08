#/bin/bash

cd "$(dirname "$0")"
cd $GOPATH/src/github.com/go-xorm/cmd/xorm/templates/goxorm
xorm reverse mysql "root:password@tcp(localhost:3306)/test" .
cd -
mv $GOPATH/src/github.com/go-xorm/cmd/xorm/templates/goxorm/model/* models-xorm

