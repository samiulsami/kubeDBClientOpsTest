package main

import (
	"ops-center/kubeDBClientOpsTest/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
