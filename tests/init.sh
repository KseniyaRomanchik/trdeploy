#!/bin/bash

BASE_PATH="${PWD}"
TEST_RES="${BASE_PATH}/example/share_resources/res1"
HCL_CONFIG="${TEST_RES}/${TERRAGRUNT_CONFIG_NAME}"
TF_FILES="${TEST_RES}/${TERRAFORM_DIR}"

cd ${TEST_RES}

${BASE_PATH}/cmd/trdeploy init --work-profile=work --prefix=trdeploy --global-var-path="${BASE_PATH}/example/var" --global-pipline-profile="${BASE_PATH}/example/pipline_profile"  --base-path="${BASE_PATH}/example" --s3-state-bucket="trdeploy-terraform-state-backet" --dynamodb-lock-table="trdeploy-terraform-state-backet-lock" --audit-profile="audit" 2> /dev/null


printf "\n\n\n"


if [ $? -eq 0 ]
then
  printf "*** Successfull init\n"
else
  printf "\033[31m Init error\n" >&2
  exit 1
fi

if [ -f "${HCL_CONFIG}" ];
then
    printf "*** ${HCL_CONFIG} exist"
else
  printf "\033[31m TERRAGRUNT_CONFIG_NAME=${TERRAGRUNT_CONFIG_NAME} was not created\n"
  exit 1
fi

if [ "${TF_FILES}" ];
then
    printf "*** ${TF_FILES} exist"
else
  printf "\033[31m TERRAGRUNT_CONFIG_NAME=${TERRAFORM_DIR} was not created\n"
  exit 1
fi