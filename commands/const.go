package commands

import (
	"log"
	"os"
	"path/filepath"
)

const (
	terragruntConfigName  = "terragrunt.hcl"
	terragruntConfigTempl = `remote_state {
	backend = "s3"
		config =  {
			bucket         = "%s"
			key            = "%s"
			region         = "%s"
			encrypt        = true
			dynamodb_table = "%s"
			profile = "%s"
		}
}`
	Init       = "init"
	Plan       = "plan"
	Apply      = "apply"
	Destroy    = "destroy"
	PipeDeploy = "pipe-deploy"

	configFileName = "config"
)

func currentPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return currentPath
}

func CurrentDir() string {
	return filepath.Base(currentPath())
}
