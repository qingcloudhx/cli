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
