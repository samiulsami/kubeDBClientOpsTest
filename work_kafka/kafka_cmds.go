package work_kafka

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var KafkaCmdTest = &cobra.Command{
	Use:   "kafkaTest",
	Short: "This is a simple CLI application",
	Long:  `Test kafka all`,
	Run: func(cmd *cobra.Command, args []string) {
		klog.Info("TestKafkaStatus():\n")
		TestKafkaStatus()
	},
}
