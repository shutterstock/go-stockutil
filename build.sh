#!/bin/bash -e

PROJECT="${PROJECT:-$(basename $PWD)}"
ORG_PATH="github.com/shutterstock"
REPO_PATH="${ORG_PATH}/${PROJECT}"

export GOPATH=${PWD}/gopath
export PATH="$GOPATH/bin:$PATH"

rm -f $GOPATH/src/${REPO_PATH}
mkdir -p $GOPATH/src/${ORG_PATH}
ln -s ${PWD} $GOPATH/src/${REPO_PATH}

eval $(go env)

go get code.google.com/p/go.tools/cmd/cover

if [ -s DEPENDENCIES ]; then
  for d in $(cat DEPENDENCIES); do
    go get $d
  done
fi

# set flags
[ "$DEBUG" == 'true' ] || GOFLAGS="-ldflags '-s'"

# build it!
for pkg in *util; do
  if [ -d $pkg ]; then
    case $1 in
    test)
      CGO_ENABLED=0 go test -test.v -coverprofile profile.file -a $GOFLAGS ${REPO_PATH}/$pkg
      go tool cover -html=profile.file -o coverage.html
      ;;
    *)
      CGO_ENABLED=0 go build -a $GOFLAGS ${REPO_PATH}/$pkg
      ;;
    esac
  fi
done
