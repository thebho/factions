#!/bin/bash
cd "$(dirname $0)"

echo "--------- Stopping MySQL container ---------"
docker-compose stop mysql

echo "--------- Removing MySQL container ---------"
docker-compose rm -f mysql

echo "--------- Building MySQL container ---------"
docker-compose up -d

sleep 5
echo "--------- Injecting Factions Data ---------"
docker exec -i mysql mysql -uroot -ppass < data/factionsdata.sql
