#!/usr/bin/env bash

# REQ: Synthesizes the app into terraform configuration JSON. <rbt 2023-10-04>

set +o braceexpand
set -o errexit
set -o noclobber
set -o noglob
set +o nounset
set -o pipefail
set -o xtrace

path=$(realpath "$0")
dir=$(dirname "$path")
cd "$dir/../"

export XENIA_NAME='xenia'
export XENIA_REGION_PRIMARY='USEast2'
export XENIA_REGION_SECONDARY='USWest2'
[[ "${XENIA_WHITELIST_IPS@a}" == *x* ]] || export XENIA_WHITELIST_IPS=''

npx cdktf synth 
