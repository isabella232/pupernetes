#!/usr/bin/env bash

set -exo pipefail

cd $(dirname $0)

curl -L https://github.com/heptio/sonobuoy/releases/download/v0.16.5/sonobuoy_0.16.5_linux_amd64.tar.gz -o sonobuoy.tar.gz
tar -xzvf sonobuoy.tar.gz
rm -v sonobuoy.tar.gz

set +e

./sonobuoy run --mode Quick --skip-preflight || exit $?

until ./sonobuoy status
do
    sleep 10
done

while true
do
    SSTATUS=$(./sonobuoy status)
    echo ${SSTATUS} | grep -c "Sonobuoy has completed" && break
    sleep 10
done
set -e

./sonobuoy status

until ls -l *_sonobuoy_*.tar.gz
do
    sleep 2
    ./sonobuoy status
    ./sonobuoy retrieve
done

./sonobuoy e2e *_sonobuoy_*.tar.gz
