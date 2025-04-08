package cmd

import (
	"ops-center/kubeDBClientOpsTest/work_kafka"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var RootCmd = &cobra.Command{
	Use:   "app",
	Short: "This is a simple CLI application",
	Long:  `A simple CLI application built with Cobra in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		klog.Infof("Run: %s", cmd.Name())
	},
}

func init() {
	RootCmd.AddCommand(work_kafka.KafkaCmdTest)
}
