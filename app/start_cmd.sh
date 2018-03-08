#!/bin/bash

sh create_models.sh
go test -bench SelectOne -count 2 -benchtime 1s -benchmem
go test -bench InsertOne -count 2 -benchtime 1s -benchmem
go test -bench InsertLarge -count 2 -benchtime 1s -benchmem
go test -bench InsertBulk -count 2 -benchtime 1s -benchmem
go test -bench SelectLarge -count 2 -benchtime 1s -benchmem
go test -bench SelectMany -count 2 -benchtime 1s -benchmem

