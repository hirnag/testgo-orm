#!/bin/sh
interval=$1
# インターバルの単位は秒で初期値は1秒
test -z $interval && interval=1
watch -n $interval --differences "mysql -uroot -Dpassword --table -Be 'show processlist;'"


