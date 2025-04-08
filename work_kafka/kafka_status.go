package work_kafka

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	utils "ops-center/kubeDBClientOpsTest/utils"

	"kubedb.dev/db-client-go/kafka"

	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1"
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

	for _, v := range brokers {
		klog.Infof("broker: %s\n", v)
	}

	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "kubedb.com",
		Version: "v1",
		Kind:    "Kafka",
	})
	if err := kbClient.Get(context.TODO(), client.ObjectKey{
		Namespace: "demo",
		Name:      "kafka",
	}, obj); err != nil {
		klog.Errorf("failed to get unstructured object: %v", err)
		return
	}

	kafkaDB := &dbapi.Kafka{}
	if err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.UnstructuredContent(), kafkaDB); err != nil {
		klog.Errorf("failed to convert unstructured object to a concrete type: %v", err)
		return
	}

	kafkaClient, err := kafka.NewKubeDBClientBuilder(kbClient, kafkaDB).
		WithContext(context.TODO()).
		GetKafkaClient()
	if err != nil {
		klog.Errorf("failed to create sarama client: %v", err)
		return
	}

	underReplicatedPartitions, err := getKafkaUnderReplicatedPartitions(kafkaClient)
	if err != nil {
		klog.Errorf("failed to get under replicated partitions: %v", err)
		return
	}

	for _, v := range underReplicatedPartitions {
		klog.Infof("under replicated partition: %s\n", v)
	}

	klog.Infof("db: %v\n", kafkaDB)
}
