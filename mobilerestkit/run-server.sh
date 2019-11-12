#!/bin/bash

mysql_host="$1"
mysql_port="$2"
timeout="$3"

until nc -w $timeout -z $mysql_host $mysql_port; do
    echo "Connection to ${mysql_host}:${mysql_port} was failed"
    sleep 1
done

echo "${mysql_host}:${mysql_port} is up - executing command"

realize start