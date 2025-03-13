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

	klog.Infof("mssql.Status.State: %s", &mssql.Status.Conditions)
	totalMemory, err := GetTotalMemoryMSSQLServer(mssql)
	if err != nil {
		klog.Errorf("error fetching total memory limit from mssqlserver: %w", err)
		return
	}
	klog.Infof("mssql memory limit: %d", *totalMemory)
	klog.Info("=======Test mssqlserver status=======")
}
