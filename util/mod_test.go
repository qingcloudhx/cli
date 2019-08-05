package util

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

/**
* @Author: hexing
* @Date: 19-7-18 下午4:33
 */
func TestExecCmd(t *testing.T) {
	cmd := exec.Command("ls")
	err := ExecCmd(cmd, "/home/code/flowgoEx/src/cli/util")
	assert.Nil(t, err)
}
func TestRun(t *testing.T) {
	cmd := exec.Command("/home/code/test/device", "-conf", "/home/code/test/ficti_flow.json")
	cmd.Env = []string{"FLOGO_LOG_LEVEL=DEBUG"}
	SetVerbose(true)
	err := ExecCmd(cmd, "/home/code/flowgoEx/src/cli/util")
	assert.Nil(t, err)
}
