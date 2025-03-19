package main

import (
	"strings"

	"github.com/easeflowHQ/flohive-node-plugin/types"
)

func main() {}

type SignatureScanDemoPlugin struct {
}

var signatures = [...]string{"flo", "FLO", "Flo"}

func (plugin *SignatureScanDemoPlugin) GetSupportedJobTypes() []string {
	return []string{"signature-scan-demo"}
}

func (plugin *SignatureScanDemoPlugin) ExecuteJob(job *types.Job) ([]byte, error) {
	memory := string(job.Task)
	for _, sig := range signatures {
		if strings.Contains(memory, sig) {
			return []byte{1}, nil
		}
	}
	return []byte{0}, nil
}

var Plugin SignatureScanDemoPlugin
