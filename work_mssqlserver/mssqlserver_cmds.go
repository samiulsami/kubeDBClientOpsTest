package work_mssqlserver

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var MSSQLCmdTest2 = &cobra.Command{
	Use:   "mssqlTest",
	Short: "This is a simple CLI application",
	Long:  `Test mssqlserver all`,
	Run: func(cmd *cobra.Command, args []string) {
		klog.Info("TestMssqlServerStatus():\n")
		TestMSSQLServerStatus()
	},
}
