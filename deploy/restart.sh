#!/bin/bash

DEPLOY_PWD=$PWD
OUTPUT_DIR=$PWD/../output

function restart() {
  set -x
  cd $DEPLOY_PWD/.. || exit
  bash $DEPLOY_PWD/../build.sh

  # cp 会直接覆盖旧的
  echo "rebuild and restart:[openapi.bittrace.proj]"
  docker cp ${OUTPUT_DIR}/openapi-cli "openapi.bittrace.proj":/bittrace/
  docker restart "openapi.bittrace.proj"

  cd $DEPLOY_PWD || exit
}

restart
