#!/bin/bash

docker cp ./resources/dump mongo:/tmp
docker exec mongo bash -c "mongorestore /tmp/dump/"


