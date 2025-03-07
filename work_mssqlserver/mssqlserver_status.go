package work_mssqlserver

import (
	"k8s.io/klog/v2"
)

func TestMSSQLServerStatus() {
	_, mssql, err := GetMSSQLServerDBAndClient()
	if err != nil {
		klog.Error(err, "failed to get mssqlserver db and client")
		return
	}

	klog.Infof("mssql.Status.State: %s", mssql)
	klog.Info("=======Test mssqlserver status=======")
}
