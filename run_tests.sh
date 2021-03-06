#/bin/bash

set -exuo pipefail

root=$PWD
arr=$(find . | grep '99_hw$' | grep -v 'ci_cd/99_hw$' | grep -v '99_hw/code' |  grep -v 'conf_monitoring/99_hw')
for i in $arr; do golangci-lint -c .golangci.yml run $i/...;done

if [ -d "$root/09_conf_monitoring/99_hw/server" ]
then
    cd $root/09_conf_monitoring/99_hw/server
    golangci-lint -c $root/.golangci.yml run ./...
fi