#!/usr/bin/env bash

function publish {
  version=$1

  if [[ -z $version ]]; then
    echo "Must provide version to publish."
    exit 1
  fi

  go mod tidy \
  && go test ./tests/... \
  && git commit -m "version/main: ${version}" \
  && git tag ${version} \
  && git push origin ${version} \
  && GOPROXY=proxy.golang.org go list -m github.com/bnert@${version}
}

if [[ -z $1 || "$1" == "help" ]]; then
  printf "Available commands:
  publish"
else
  cmd=$1
  if [[ "$cmd" == "publish" ]]; then
    shift;
    publish $@
  fi
fi
