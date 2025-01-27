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

	memory, err := GetTotalMemoryMSSQLServer(mssql)
	if err != nil {
		klog.Error(err, "failed to get total memory of mssqlserver")
		return
	}

	klog.Infof("total memory of mssqlserver: %d", *memory)
	klog.Info("=======Test mssqlserver status=======")
}
