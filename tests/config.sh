#!/bin/bash

set -a
source .env
set +a

BASE_PATH="${PWD}"
TEST_RES="${BASE_PATH}/example/share_resources/res1"

rm -fr "${HOME}/.${CONFIG_FILE_NAME}.yaml"
cd ${TEST_RES}

printf "Start testing apply with flags in ${TEST_RES}...\n"

${BASE_PATH}/cmd/trdeploy apply --work-profile=work --prefix=trdeploy --global-var-path="${BASE_PATH}/example/var" --global-pipeline-profile="${BASE_PATH}/example/pipline_profile" --base-path="${BASE_PATH}/example" --s3-state-bucket="trdeploy-terraform-state-backet" --dynamodb-lock-table="trdeploy-terraform-state-backet-lock" --audit-profile=audit 2> /dev/null

if [ $? -eq 0 ];
then
  printf "*** Successfully apply with flags\n\n"
else
  printf "\033[31m apply with flags error\n" >&2
  exit 1
fi


echo "work:
  global-var-path: ${BASE_PATH}/example/var
  global-pipeline-profile: ${BASE_PATH}/example/pipeline_profile
  base-path: ${BASE_PATH}/example
  s3-state-bucket: trdeploy-terraform-state-backet
  dynamodb-lock-table: trdeploy-terraform-state-backet-lock
  audit-profile: audit" > "${HOME}/.${CONFIG_FILE_NAME}.yaml"

printf "Start testing apply with loading from config in ${TEST_RES}...\n"

${BASE_PATH}/cmd/trdeploy apply --work-profile=work --prefix=trdeploy 3> /dev/null

if [ $? -eq 0 ];
then
  printf "*** Successfully load from ${CONFIG_FILE_NAME}\n"
else
  printf "\033[31m Loading from ${CONFIG_FILE_NAME} error\n" >&3
  exit 1
fi