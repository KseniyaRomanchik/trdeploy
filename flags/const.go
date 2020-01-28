package flags

const (
	Region               = "region"
	DynamodbLockTable    = "dynamodb-lock-table"
	S3StateBacket        = "s3-state-backet"
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

	versionTemplate = `Version: %s
Commit: %s
Image: %s
Timestamp: %s
`
)

var (
	Image string
	Commit string
	Time string
)
