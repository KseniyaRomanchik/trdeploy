package flags

const (
	Region               = "region"
	DynamoDBLockTable    = "dynamodb-lock-table"
	S3StateBucket        = "s3-state-bucket"
	AuditProfile         = "audit-profile"
	OutPlanLog           = "out-plan-log"
	AdditionalArgs       = "additional-args"
	ModuleTfvars         = "module-tfvars"
	GlobalVarPath        = "global-var-path"
	GlobalPiplineProfile = "global-pipline-profile"
	BasePath             = "base-path"
	Config               = "config"
	PlanFile             = "plan-file"
	DeployProfile        = "deploy-profile"
	Prefix               = "prefix"
	WorkProfile          = "work-profile"
	Multithread          = "multithread"
	PiplineFile          = "pipeline-file"
	Parallelism          = "parallelism"
	Timeout              = "timeout"
	Test                 = "test"

	versionTemplate = `Version: %s
Commit: %s
Image: %s
Timestamp: %s
`
)

var (
	Image  string
	Commit string
	Time   string
)
