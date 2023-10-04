#!/usr/bin/env bash

# REQ: Generates CDK constructs. <rbt 2023-09-29>

set +o braceexpand
set -o errexit
set -o noclobber
set -o noglob
set -o nounset
set -o pipefail
set -o xtrace

path=$(realpath "$0")
dir=$(dirname "$path")
cd "$dir"

npm --version
npm cache verify
npm install

npx cdktf --version
npx cdktf get
