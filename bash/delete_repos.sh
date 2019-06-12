#!/bin/bash

function git_repo_delete(){
  curl -vL \
    -H "Authorization: token $GITHUB_SECRET" \
    -H "Content-Type: application/json" \
    -X DELETE https://api.github.com/repos/$2/$1 \
    | jq .
}

repos=$(cat $1)
for repo in $repos; do (git_repo_delete "$repo" "$2"); done
