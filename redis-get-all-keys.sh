#!/bin/bash

for id in {1..4}
do
  curl -i "http://localhost:8000/get/$id"
done


