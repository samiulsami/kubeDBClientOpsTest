package work_kafka

import (
	"context"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func getKafkaBrokers(ctx context.Context, ctrlClient client.Client, name, namespace string) ([]string, error) {
	appBinding := &unstructured.Unstructured{}
	appBinding.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "appcatalog.appscode.com",
		Version: "v1alpha1",
		Kind:    "AppBinding",
	})

	if err := ctrlClient.Get(
		ctx,
		client.ObjectKey{
			Namespace: namespace,
			Name:      name,
		},
		appBinding,
	); err != nil {
		return nil, fmt.Errorf("failed to get appbinding: %w", err)
	}

	brokers, found, err := unstructured.NestedString(appBinding.Object, "spec", "clientConfig", "url")
	if err != nil {
		return nil, fmt.Errorf("failed to get brokers: %w", err)
	}
	if !found {
		return nil, fmt.Errorf("brokers not found")
	}

	return strings.Split(brokers, ","), nil
}

func getKafkaUnderReplicatedPartitions(client sarama.Client) ([]string, error) {
	topics, err := client.Topics()
	if err != nil {
		return nil, fmt.Errorf("failed to get topics: %w", err)
	}

	admin, err := sarama.NewClusterAdminFromClient(client)
	if err != nil {
		return nil, fmt.Errorf("failed to create admin client: %w", err)
	}
	defer admin.Close()

	metadata, err := admin.DescribeTopics(topics)
	if err != nil {
		return nil, fmt.Errorf("error describe topics: %w", err)
	}

	underReplicatedPartitions := make([]string, 0)
	for _, topicMetadata := range metadata {
		for _, partitionMetadata := range topicMetadata.Partitions {
			replicas := len(partitionMetadata.Replicas)
			inSyncReplicas := len(partitionMetadata.Isr)

			if inSyncReplicas < replicas {
				underReplicatedPartitions = append(
					underReplicatedPartitions,
					fmt.Sprintf(
						"%s:%d",
						topicMetadata.Name,
						partitionMetadata.ID,
					),
				)
			}
		}
	}

	return underReplicatedPartitions, nil
}
