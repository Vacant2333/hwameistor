package cmdparser

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/hwameistor/hwameistor/pkg/hwameictl/cmdparser/definitions"
	"github.com/hwameistor/hwameistor/pkg/hwameictl/cmdparser/volume"
)

var Hwameictl = &cobra.Command{
	Use:   "hwameictl",
	Short: "Hwameictl is the command-line tool for Hwameistor.",
	Long: "Hwameictl is a tool that can manage all Hwameistor resources and their entire lifecycle.\n" +
		"Complete documentation is available at https://hwameistor.io/",
	RunE: func(cmd *cobra.Command, args []string) error {
		// root cmd will show help only
		return cmd.Help()
	},
}

func init() {
	// Hwameictl flags
	Hwameictl.PersistentFlags().StringVar(&definitions.Kubeconfig, "kubeconfig", "~/.kube/config", "Specify the kubeconfig file")
	Hwameictl.PersistentFlags().DurationVar(&definitions.Timeout, "timeout", 3*time.Second, "Set the request timeout")

	// Sub command
	Hwameictl.AddCommand(volume.Volume)

	// todo: add debug flag
	log.SetLevel(log.PanicLevel)
}