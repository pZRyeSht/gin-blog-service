#!/usr/bin/env bash

ROOT=`pwd`

go run $ROOT/cmd/mysql-db-migrate/main.go -conf $ROOT/configs/ -do drop
go run $ROOT/cmd/mysql-db-migrate/main.go -conf $ROOT/configs/ -do create
go run $ROOT/cmd/mysql-db-migrate/main.go -conf $ROOT/configs/ -do migrate