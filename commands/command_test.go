package commands

import (
	"flag"
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var command = Command{
	Cmd:        &exec.Cmd{},
	ThreadName: []string{},
}

func TestDir(t *testing.T) {
	t.Parallel()
	dir := "testDir"
	command = Dir(dir)(command)

	if command.Dir != dir {
		t.Errorf("Invalid cmd Dir(): %+v != %+v", command.Dir, dir)
	}
}

func TestThreadName(t *testing.T) {
	t.Parallel()
	thNames := []string{"name1", "name2"}
	command = ThreadName(thNames...)(command)

	if !reflect.DeepEqual(command.ThreadName, thNames) {
		t.Errorf("Invalid cmd ThreadName(): %+v != %+v", command.ThreadName, thNames)
	}
}

func TestCmd(t *testing.T) {
	t.Parallel()
	ctx := cli.NewContext(&cli.App{}, &flag.FlagSet{}, nil)

	argString := strings.Join([]string{"-var-file", fmt.Sprintf("%s/common.tfvars", ""),
		"-var-file", fmt.Sprintf("%s/%s.tfvars", "", ""),
		"-var-file", "",
		"-var", fmt.Sprintf("prefix=%s", ""),
		"-var", fmt.Sprintf("aws_audit=%s", "")}, " ")

	command = Cmd(ctx)(command)

	if !strings.Contains(command.String(), argString) {
		t.Errorf("Invalid cmd Cmd(): \n%+v not contains \n%+v", command.String(), argString)
	}
}

func TestAutoApprove(t *testing.T) {
	t.Parallel()
	flag := "-auto-approve"
	command = AutoApprove()(command)

	for _, arg := range command.Args {
		if arg == flag {
			return
		}
	}

	t.Errorf("Invalid cmd AutoApprove(): %+v - %+v", command.Args, flag)
}

func TestParallelism(t *testing.T) {
	t.Parallel()
	flag := "-parallelism"
	value := 1
	command = Parallelism(value)(command)

	for i, arg := range command.Args {
		if arg == flag && command.Args[i + 1] == strconv.Itoa(value) {
			return
		}
	}

	t.Errorf("Invalid cmd Parallelism(): %+v - %+v=%+v", command.Args, flag, value)
}
