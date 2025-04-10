package work_kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func getKafkaBrokers(ctx context.Context, ctrlClient client.Client, name, namespace string) (string, error) {
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
		return "", fmt.Errorf("failed to get appbinding: %w", err)
	}

	brokers, found, err := unstructured.NestedString(appBinding.Object, "spec", "clientConfig", "url")
	if err != nil {
		return "", fmt.Errorf("failed to get brokers: %w", err)
	}
	if !found {
		return "", fmt.Errorf("brokers not found")
	}

	return brokers, nil
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

func verifyKafkaBrokers(client sarama.Client, brokers []string) ([]byte, error) {
	if client == nil {
		return nil, fmt.Errorf("client is nil")
	}

	clusterBrokers := client.Brokers()
	if len(clusterBrokers) != len(brokers) {
		return nil, fmt.Errorf("expected '%d' brokers, but found '%d' ", len(brokers), len(clusterBrokers))
	}

	brokersStatus := []byte{}
	for _, broker := range clusterBrokers {
		if err := broker.Open(client.Config()); err != nil {
			brokersStatus = fmt.Appendf(
				brokersStatus,
				"Broker '%s' is down: '%v'\n",
				broker.Addr(),
				err,
			)
			continue
		}

		connected, err := broker.Connected()
		if err != nil {
			brokersStatus = fmt.Appendf(
				brokersStatus,
				"Failed to check connection status for broker '%s': %v\n",
				broker.Addr(),
				err,
			)
			continue
		}

		if !connected {
			brokersStatus = fmt.Appendf(
				brokersStatus,
				"Broker '%s' is down\n",
				broker.Addr(),
			)
			continue
		}

		brokersStatus = fmt.Appendf(
			brokersStatus,
			"Broker '%s' is up\n",
			broker.Addr(),
		)

		if err := broker.Close(); err != nil {
			brokersStatus = fmt.Appendf(
				brokersStatus,
				"Failed to close broker '%s': %v\n",
				broker.Addr(),
				err,
			)
		}
	}

	return brokersStatus, nil
}
