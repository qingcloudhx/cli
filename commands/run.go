package commands

import (
	"fmt"
	"github.com/qingcloudhx/cli/api"
	"github.com/qingcloudhx/cli/common"
	"github.com/qingcloudhx/cli/util"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

/**
* @Author: hexing
* @Date: 19-8-5 上午11:20
 */
var rflogoJsonPath string

func init() {
	RunCmd.Flags().StringVarP(&rflogoJsonPath, "file", "f", "", "specify a flogo.json to create project from")
	rootCmd.AddCommand(RunCmd)
}

var RunCmd = &cobra.Command{
	Use:              "run [flags] [appName]",
	Short:            "run a flogo application project",
	Long:             `run a flogo application project.`,
	Args:             cobra.RangeArgs(0, 1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {

		//If a jsonFile is specified in the build.
		//Create a new project in the temp folder and copy the bin.
		util.SetVerbose(true)
		tempDir, err := api.GetTempDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting temp dir: %v\n", err)
			os.Exit(1)
		}
		defer func() {
			os.RemoveAll(tempDir)
		}()
		tempProject, err := api.CreateProject(tempDir, "", rflogoJsonPath, "latest")
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(os.Stderr, "Error creating temp project tempDir:%s,err: %+v\n", tempDir, err)
			os.Exit(1)
		}

		common.SetCurrentProject(tempProject)

		options := api.BuildOptions{Shim: "", OptimizeImports: false, EmbedConfig: true}

		err = api.BuildProject(common.CurrentProject(), options)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error building temp project: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("exec:%s,dir:%s,temp:%s\n", tempProject.Executable(), tempProject.BinDir(), tempDir)
		cx := exec.Command(tempProject.Executable())
		cx.Env = []string{"FLOGO_LOG_LEVEL=DEBUG"}
		err = util.ExecCmd(cx, tempProject.BinDir())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error run temp project: %v\n", err)
			os.Exit(1)
		}
	},
}
