#!/bin/bash
go build -o main main.go
go build -o main2 main2.go
echo "Print 1M args in cli"
echo -e "\nMain with loop\n"
time ./main `seq 1 100000` > /dev/null
echo -e "\nMain with strings join\n"
time ./main2 `seq 1 100000` > /dev/null