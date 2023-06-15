#!/bin/bash

for id in {1..3}
do
  curl -i "http://localhost:8000/get/$id"
done


