#!/bin/sh

./transfer.sh ./trainData train.csv
./transfer.sh ./testData test.csv

go run knn.go
