#!/usr/bin/env bash

# set The set command enables options within a script
# -o At the point in the script where you want the options to take effect, use set -o option-name 
# -e Exit immediately if a pipeline returns a non-zero status
# -u Treat unset variables and parameters (except ‘@’ or ‘*’) as an error when performing parameter expansion.

set -euo

# for periodics maybe this is git trickery, or maybe this is gsutil cat gs://some/well/know/path

# ':-""' means if PULL_BASE_SHA is not set already then set it to ""
BASE_SHA="${PULL_BASE_SHA:-""}"

# I assume artifacts are just temp files created in scripts?
ARTIFACTS="${ARTIFACTS:-tmp/artifacts}"

# With -p you can create sub-directories of a directory
mkdir -p "${ARTIFACTS}"

# will this work here or was this only for scripts in verify dir?
# and does this also install go or do we need to do it?
kube::golang::verify_go_version

# Explicitly opt into go modules, even though we're inside a GOPATH directory
export GO111MODULE=on

# install depstat
go install github.com/kubernetes-sigs/depstat

function write_report() { depstat stats --json }
function diff_reports() { diff $1 $2 }

write_report > "${ARTIFACTS}/stats.json"

# “-n” flag returns true if a string is not null
if [ -n "${BASE_SHA}" ]; then
  git checkout -b base "${BASE_SHA}"
  write_report > "${ARTIFACTS}/stats-base.json"
  git checkout HEAD@{1}
  # is this some way of specifying two files together as arguments to the diff_reports function?
  diff_reports "${ARTIFACTS}"/report{-base}.json  
fi