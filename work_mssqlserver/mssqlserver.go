package work_mssqlserver

import (
	"fmt"

	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kmapi "kmodules.xyz/client-go/api/v1"
	dbapiv1alpha2 "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r RemedyCheck) getMSSQLServerDB(appRef *core.TypedObjectReference) (*dbapiv1alpha2.MSSQLServer, error) {
	ref := kmapi.ObjectReference{
		Name:      appRef.Name,
		Namespace: *appRef.Namespace,
	}

	version, err := findResourceVersion(r.restMapper, schema.GroupKind{
		Group: *appRef.APIGroup,
		Kind:  appRef.Kind,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find version : %v", err)
	}

	obj, err := r.getK8sObject(schema.GroupVersionKind{
		Group:   *appRef.APIGroup,
		Kind:    appRef.Kind,
		Version: version,
	}, ref)
	if err != nil {
		return nil, fmt.Errorf("failed to get k8s object : %v", err)
	}

	db := &dbapiv1alpha2.MSSQLServer{}
	err = runtime.DefaultUnstructuredConverter.
		FromUnstructured(obj.UnstructuredContent(), db)
	if err != nil {
		return nil, fmt.Errorf("failed to convert unstructured object to a concrete type: %w", err)
	}

	return db, nil
}

func GetMSSQLServerDB(kbClient client.Client) (*dbapi.Postgres, error) {
	ref := kmapi.ObjectReference{
		Name:      "mssqlserver",
		Namespace: "demo",
	}
	gvk := schema.GroupVersionKind{
		Version: "v1alpha2",
		Group:   "kubedb.com",
		Kind:    "MSSQLServer",
	}

	obj, err := utils.GetK8sObject(gvk, ref, kbClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get k8s object : %v", err)
	}

	db := &dbapi.Postgres{}
	err = runtime.DefaultUnstructuredConverter.
		FromUnstructured(obj.UnstructuredContent(), db)
	if err != nil {
		return nil, fmt.Errorf("failed to convert unstructured object to a concrete type: %w", err)
	}

	return db, nil
}

// func (r RemedyCheck) getMSSQLServerClient(db *dbapiv1alpha2.MSSQLServer) (*mssqlserver.Client, error) {
// 	kubeDBClient, err := mssql.NewKubeDBClientBuilder(r.ctrlClient, db).
// 		WithContext(context.Background()).
//
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to build kubedb mssqlserver client : %v", err)
// 	}
//
// 	return kubeDBClient, nil
// }
