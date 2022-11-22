#! bin/bash

cd /app/src

go mod tidy
air >> .air.log
