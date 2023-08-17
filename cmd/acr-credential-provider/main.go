package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/kuromesi/acr-credential-provider/pkg/acr"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := newCredentialProviderCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

var gitVersion string

func newCredentialProviderCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "acr-credential-provider",
		Short:   "ACR credential provider for kubelet",
		Version: gitVersion,
		Run: func(cmd *cobra.Command, args []string) {
			p := acr.NewCredentialProvider(&acr.Client{})
			if err := p.Run(context.TODO()); err != nil {
				klog.Errorf("Error running credential provider plugin: %v", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}
