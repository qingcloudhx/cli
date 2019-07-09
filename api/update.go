package api

import (
	"fmt"
	"os/exec"

	"github.com/qingcloudhx/cli/common"
	"github.com/qingcloudhx/cli/util"
)

func UpdatePkg(project common.AppProject, pkg string) error {

	if Verbose() {
		fmt.Printf("Updating Package: %s", pkg)
	}

	err := util.ExecCmd(exec.Command("go", "get", "-u", pkg), project.SrcDir())
	return err
}
