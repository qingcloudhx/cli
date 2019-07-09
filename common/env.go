package common

import (
	"github.com/qingcloudhx/cli/util"
)

var verbose = false
var appProject AppProject

func SetVerbose(enable bool) {
	verbose = enable
	util.SetVerbose(enable)
}

func Verbose() bool {
	return verbose
}

func CurrentProject() AppProject {
	return appProject
}

func SetCurrentProject(project AppProject) {
	appProject = project
}
