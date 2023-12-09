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

export XEN_NAME='xen'
export XEN_REGION_PRIMARY='USEast2'
export XEN_REGION_SECONDARY='USWest2'
[[ "${XEN_WHITELIST_IPS@a}" == *x* ]] || export XEN_WHITELIST_IPS=''

mkdir -p './out'
npx cdktf synth 
