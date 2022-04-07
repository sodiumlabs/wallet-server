#!/bin/bash
set -o errexit

### get project dir
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do
  DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE"
done
DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
readonly PROJECT_ROOT="$(dirname $DIR)"
RUN_ROOT="$PROJECT_ROOT"
cd $PROJECT_ROOT;


projects=$(find "${PROJECT_ROOT}/abis" -type file | xargs -I {} basename {});
packages="abis";

for project in $projects
do
    if [[ -f "${PROJECT_ROOT}/${packages}/${project}" ]];
    then
        filename="${project%.*}"
        abigen --abi "${PROJECT_ROOT}/${packages}/${project}" --pkg contracts --type $filename --out "${PROJECT_ROOT}/pkg/contracts/${filename,}.go"
    fi;
done