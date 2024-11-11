package cmd

import (
	"fmt"

	utils "github.com/shn27/Test/utils"
	"github.com/shn27/Test/work"
	"github.com/shn27/Test/work_postgres"
	"github.com/spf13/cobra"
	kmapi "kmodules.xyz/client-go/api/v1"
)

var runbook = &cobra.Command{
	Use:   "runbook",
	Short: "Greet the user",
	Long:  `This subcommand greets the user with a custom message.`,
	Run: func(cmd *cobra.Command, args []string) {
		work.IsRunbookCRExit()
	},
}

var markdown1 = &cobra.Command{
	Use:   "markdown",
	Short: "Greet the user",
	Long:  `This subcommand greets the user with a custom message.`,
	Run: func(cmd *cobra.Command, args []string) {
		work.GetMarkdown()
	},
}

var tableWriter = &cobra.Command{
	Use:   "table",
	Short: "Greet the user",
	Long:  `This subcommand greets the user with a custom message.`,
	Run: func(cmd *cobra.Command, args []string) {
		work.Table()
	},
}

var version = &cobra.Command{
	Use:   "version",
	Short: "Greet the user",
	Long:  `This subcommand greets the user with a custom message.`,
	Run: func(cmd *cobra.Command, args []string) {
		ref := kmapi.TypedObjectReference{
			APIGroup: "kubedb.com", // k8s_group
			Kind:     "MongoDB",    // k8s_kind
		}
		version, err := work.GetPreferredResourceVersion(ref)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("=====================", version)
	},
}

var resource = &cobra.Command{
	Use:   "resource",
	Short: "Greet the user",
	Long:  `This subcommand greets the user with a custom message.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.GetResource("mongodb-down")
	},
}

var currentop = &cobra.Command{
	Use:   "currentop",
	Short: "Greet the user",
	Long:  `This subcommand greets the user with a custom message.`,
	Run: func(cmd *cobra.Command, args []string) {
		work.DbCurrentOp()
	},
}

var mongodbServerStatus = &cobra.Command{
	Use:   "mongodbServerStatus",
	Short: "Greet the user",
	Long:  `This subcommand greets the user with a custom message.`,
	Run: func(cmd *cobra.Command, args []string) {
		work.Network()
	},
}

var mongoBDMetricsCursorOpen = &cobra.Command{
	Use:   "cursor",
	Short: "Greet the user",
	Long:  `This subcommand greets the user with a custom message.`,
	Run: func(cmd *cobra.Command, args []string) {
		work.Network()
	},
}

var RootCmd = &cobra.Command{
	Use:   "app",
	Short: "This is a simple CLI application",
	Long:  `A simple CLI application built with Cobra in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		work.Ans()
	},
}

func init() {
	RootCmd.AddCommand(markdown1)
	RootCmd.AddCommand(runbook)
	RootCmd.AddCommand(version)
	RootCmd.AddCommand(resource)
	RootCmd.AddCommand(currentop)
	RootCmd.AddCommand(mongodbServerStatus)
	RootCmd.AddCommand(tableWriter)
	RootCmd.AddCommand(mongoBDMetricsCursorOpen)
	RootCmd.AddCommand(work_postgres.PgCmdTest)
	RootCmd.AddCommand(work_postgres.PgCmdTest2)
	RootCmd.AddCommand(work_postgres.PgCmdTestSharedBuffers)
}
