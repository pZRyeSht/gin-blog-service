#!/usr/bin/env bash

ROOT=`pwd`

go run $ROOT/cmd/mysql-db-migrate/main.go -conf $ROOT/config/ -do drop
go run $ROOT/cmd/mysql-db-migrate/main.go -conf $ROOT/config/ -do create
go run $ROOT/cmd/mysql-db-migrate/main.go -conf $ROOT/config/ -do migrate