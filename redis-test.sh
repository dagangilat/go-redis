#!/bin/sh

echo "Set values"
curl -X POST -H "Content-Type: application/json" -d '{"1":"Orange"}' http://localhost:8001/set
echo "1 : Orange"
curl -X POST -H "Content-Type: application/json" -d '{"2":"Apple"}' http://localhost:8001/set 
echo "2 : Apple"
curl -X POST -H "Content-Type: application/json" -d '{"3":"Banana"}' http://localhost:8001/set
echo "3 : Banana"

echo "Get values"
curl http://localhost:8001/get/1                                                              
curl http://localhost:8001/get/2
curl http://localhost:8001/get/3



