package work_postgres

import (
	"encoding/json"

	gohumanize "github.com/dustin/go-humanize"
	utils "github.com/shn27/Test/utils"
	"k8s.io/klog/v2"
)

func TestPostgresServerStatus() {
	kubeClient, err := utils.GetKBClient()
	if err != nil {
		klog.Error(err, "failed to get kube client")
		return
	}

	db, err := GetPostgresDB(kubeClient)
	if err != nil {
		klog.Error(err, "failed to get postgres db")
		return
	}

	pgClient, err := GetPostgresClient(kubeClient, db)
	if err != nil {
		klog.Error(err, "failed to get postgres client")
		return
	}

	stats := pgClient.DB.Stats()
	klog.Info("=======Test postgres server stats=======")

	prettyData, err := json.MarshalIndent(stats, "  ", "   ")
	if err != nil {
		klog.Error(err, "failed to marshal db stats")
	}

	klog.Info(string(prettyData))
}

func TestClientFuncs() {
	kubeClient, err := utils.GetKBClient()
	if err != nil {
		klog.Error(err, "failed to get kube client")
		return
	}

	db, err := GetPostgresDB(kubeClient)
	if err != nil {
		klog.Error(err, "failed to get postgres db")
		return
	}

	pgClient, err := GetPostgresClient(kubeClient, db)
	if err != nil {
		klog.Error(err, "failed to get postgres client")
		return
	}

	err = pgClient.DB.Ping()
	if err != nil {
		klog.Error(err, "failed to ping postgres")
		return
	}

	klog.Info("Pinged postgres\n")
	klog.Infof("pgClient.DB.Stats().InUse : %d", pgClient.DB.Stats().InUse)
}

func TestCheckAvailableSharedBuffers() {
	kubeClient, err := utils.GetKBClient()
	if err != nil {
		klog.Error(err, "failed to get kube client")
		return
	}

	db, err := GetPostgresDB(kubeClient)
	if err != nil {
		klog.Error(err, "failed to get postgres db")
		return
	}

	pgClient, err := GetPostgresClient(kubeClient, db)
	if err != nil {
		klog.Error(err, "failed to get postgres client")
		return
	}

	totalMemory, err := GetTotalMemory(pgClient, db)
	if err != nil {
		klog.Error(err, "failed to get total memory")
		return
	}

	sharedBuffersStr, err := GetSharedBuffers(pgClient)
	if err != nil {
		klog.Error(err, "failed to get shared buffers")
		return
	}

	sharedBuffers, err := gohumanize.ParseBytes(sharedBuffersStr)
	if err != nil {
		klog.Error(err, "failed to parse shared buffers")
		return
	}
	klog.Infof("Total memory: %s\n", gohumanize.IBytes(uint64(totalMemory)))
	klog.Infof("Shared buffers: %s\n", gohumanize.IBytes(uint64(sharedBuffers)))

	percentage := float64(sharedBuffers) / float64(totalMemory)
	klog.Infof("Shared buffers percentage: %.2f%%\n", percentage*float64(100))
}

func TestCheckEffectiveCacheSize() {
	kubeClient, err := utils.GetKBClient()
	if err != nil {
		klog.Error(err, "failed to get kube client")
		return
	}

	db, err := GetPostgresDB(kubeClient)
	if err != nil {
		klog.Error(err, "failed to get postgres db")
		return
	}

	pgClient, err := GetPostgresClient(kubeClient, db)
	if err != nil {
		klog.Error(err, "failed to get postgres client")
		return
	}

	totalMemory, err := GetTotalMemory(pgClient, db)
	if err != nil {
		klog.Error(err, "failed to get total memory")
		return
	}

	effectiveCacheSizeStr, err := GetEffectiveCacheSize(pgClient)
	if err != nil {
		klog.Error(err, "failed to get shared buffers")
		return
	}

	effectiveCacheSize, err := gohumanize.ParseBytes(effectiveCacheSizeStr)
	if err != nil {
		klog.Error(err, "failed to parse shared buffers")
		return
	}
	klog.Infof("Total memory: %s\n", gohumanize.IBytes(uint64(totalMemory)))
	klog.Infof("effective cache size: %s\n", gohumanize.IBytes(uint64(effectiveCacheSize)))

	percentage := float64(effectiveCacheSize) / float64(totalMemory)
	klog.Infof("effective cache size percentage: %.2f%%\n", percentage*float64(100))
}
