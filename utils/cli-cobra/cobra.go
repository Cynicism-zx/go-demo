package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

/*
构建命令行程序
*/

// 根命令
var hugoCmd = &cobra.Command{
	Use:   "git",
	Short: "Git is a distributed version control system.",
	Long:  `Git is a free and open source distributed version control system designed to handle everything from small to very large projects with speed and efficiency.`,
}

// 子命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version subcommand show git version info.",

	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecuteCommand("git", "version", args...)
		if err != nil {
			Error(cmd, args, err)
		}

		fmt.Fprint(os.Stdout, output)
	},
}

var optionCmd = &cobra.Command{
	Use:   "option",
	Short: "option subcommand show git option info.",

	Run: func(cmd *cobra.Command, args []string) {
		cm := exec.Command("git")
		output, err := cm.CombinedOutput()
		if err != nil {
			Error(cmd, args, err)
		}

		fmt.Fprint(os.Stdout, output)
	},
}

func init() {
	// 添加子命令
	hugoCmd.AddCommand(versionCmd)
	hugoCmd.AddCommand(optionCmd)
}

func ExecuteCommand(name string, subname string, args ...string) (string, error) {
	args = append([]string{subname}, args...)

	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()

	return string(bytes), err
}

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v\n", cmd.Name(), args, err)
	os.Exit(1)
}

func Exec() error {
	err := hugoCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
