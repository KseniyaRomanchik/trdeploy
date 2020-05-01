#!/bin/bash

set -a
source .env
set +a

BASE_PATH="${PWD}"
TEST_RES="${BASE_PATH}/example/share_resources/res1"
HCL_CONFIG="${TEST_RES}/${TERRAGRUNT_CONFIG_NAME}"
TF_FILES="${TEST_RES}/${TERRAFORM_DIR}"

cd ${TEST_RES}

printf "Start init tesing in ${TEST_RES}...\n"

${BASE_PATH}/cmd/trdeploy init --work-profile=work --prefix=trdeploy --global-var-path="${BASE_PATH}/example/var" --global-pipeline-profile="${BASE_PATH}/example/pipeline_profile"  --base-path="${BASE_PATH}/example" --s3-state-bucket="trdeploy-terraform-state-backet" --dynamodb-lock-table="trdeploy-terraform-state-backet-lock" --audit-profile="audit" 2> /dev/null

if [ $? -eq 0 ];
then
  printf "*** Successfull init in ${TEST_RES}\n"
else
  printf "\033[31m Init error\n" >&2
  exit 1
fi

if [ -f "${HCL_CONFIG}" ];
then
    printf "*** ${TERRAGRUNT_CONFIG_NAME} exist\n"
else
  printf "\033[31m ${TERRAGRUNT_CONFIG_NAME} was not created\n"
  exit 1
fi

if [ "${TF_FILES}" ];
then
    printf "*** ${TERRAFORM_DIR} exist\n"
else
  printf "\033[31m ${TERRAFORM_DIR} was not created\n"
  exit 1
fi