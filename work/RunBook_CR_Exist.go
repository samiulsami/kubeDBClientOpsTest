package work

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	utils "ops-center/kubeDBClientOpsTest/utils"
)

func IsRunbookCRExit() {
	myClient, err := utils.GetKBClient()
	if err != nil {
	}
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "monitoring.appscode.com",
		Kind:    "Runbook",
		Version: "v1alpha1",
	})
	err = myClient.Get(context.TODO(), client.ObjectKey{
		Name: "mongodb-down",
	}, obj)
	if err != nil {
		fmt.Println(err)
	}
}
