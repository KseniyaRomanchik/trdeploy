trdeploy
================================

![build](https://github.com/KseniyaRomanchik/trdeploy/workflows/trdeploy/badge.svg?branch=master&event=push)

#### Commands

`init`

`plan`

`apply`

`destroy`

`pipe deploy`

`pipe destroy`

#### Flags

`--region` aws region (default: "us-west-2")

`--dynamodb-lock-table` terraform state-lock table name (default: "unitedsoft-terraform-state-backet-lock")

`--s3-state-bucket` terraform s3 state bucket name (default: "unitedsoft-terraform-state-backet")

-`-audit-profile` aws audit profile name (default: "default")

`--out-plan-log` ???desc???

`--additional-args` additional args

`--module-tfvars` name of module tfvars-file  (default  {aws-profile}.tfvars)

`--global-var-path` ???desc??? (default from config ~/trdeploy.yaml)

`--global-pipline-profile` ???desc??? (default from ~/trdeploy.yaml)

`--base-path` ???desc??? (default from ~/tdeploy.yaml)

`--config` (default: "~/.trdeploy.yaml")

`--deploy-profile` ???desc???

`--parallelism` limit the number of concurrent terraform operation. (default: 1 for `destroy` and `pipe destroy`. [Terraform docs](https://www.terraform.io/docs/commands/apply.html#parallelism-n))

`--config` config file path (default: "~/.trdeploy.yaml", "/etc/.trdeploy.yaml")

`--timeout` command timeout in seconds (default: 600)

`--log-level` (default: 4)

`--prefix` ???desc???

`--work-profile` ???desc???

`--multithread` pipe threads parallel mode (default: false)

`--pipeline-file` pipeline file name

`--plan-file` ???desc???

`--help`, `-h` show help

`--version`, `-v` print the version, commit, image and timestamp of build

#### Usage

`trdeploy pipe deploy --work-profile=work --prefix=trdeploy --pipeline-file=test.yaml`

`trdeploy pipe destroy --work-profile=work --prefix=trdeploy --pipeline-file=test.yaml`

`trdeploy pipe deploy --work-profile=prod --prefix=default --pipeline-file=prod.yaml --multithread=true`

`trdeploy apply --work-profile=work --prefix=trdeploy`

#### Config

```yaml
work:
  global-var-path: {pipe_base_path}/example/var
  global-pipline-profile: {pipe_base_path}/example/pipline_profile
  base-path: {pipe_base_path}/example
  s3-state-bucket: terraform-state-backet
  dynamodb-lock-table: terraform-state-backet-lock
  audit-profile: audit
test:
  global-var-path: {pipe_base_path}/example/var
  global-pipline-profile: {pipe_base_path}/example/var/pipline_profile
  base-path: {pipe_base_path}/example
  s3-state-bucket: terraform-state-backet
  dynamodb-lock-table: terraform-state-backet-lock
  audit-profile: audit
```

#### Pipeline File

```yaml
steps:
  - name: step1
    threads:
      - name : res0
        path: share_resources/res0
        parallelism: 5
  - name: step2
    threads:
      - name : res1
        path: share_resources/res1
      - name : res2
        path: share_resources/res2
        parallelism: 5
  - name: step3
    threads:
      - name : res3
        path: share_resources/res3
```
