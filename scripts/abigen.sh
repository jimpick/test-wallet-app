#! /bin/bash

set -eo pipefail

mkdir -p ../tmp ../abigen

# FEVM

for c in \
  SimpleCoin \
  ; do
  echo Running abigen for: $c
  abigen \
    --abi=../out/$c.abi \
    --type=$c \
    --pkg=abigen \
    --out=../abigen/$c.go
done
