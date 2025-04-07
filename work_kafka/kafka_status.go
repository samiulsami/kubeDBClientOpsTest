package work_kafka

import (
	"context"

	"k8s.io/klog/v2"

	utils "ops-center/kubeDBClientOpsTest/utils"
)

func TestKafkaStatus() {
	kbClient, err := utils.GetKBClient()
	if err != nil {
		klog.Errorf("failed to get kube client: %v", err)
		return
	}
	brokers, err := getKafkaBrokers(context.TODO(), kbClient, "kafka", "demo")
	if err != nil {
		klog.Errorf("failed to get kafka brokers: %v", err)
		return
	}

	underReplicatedPartitions, err := getKafkaUnderReplicatedPartitions(
		context.TODO(),
		kbClient,
		brokers,
		"kafka",
		"demo",
	)
	if err != nil {
		klog.Errorf("failed to get under replicated partitions: %v", err)
		return
	}

	for _, v := range underReplicatedPartitions {
		klog.Infof("under replicated partition: %s\n", v)
	}
}
