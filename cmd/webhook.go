package cmd

import (
	"net/http"

	"github.com/jc0b/fleetdm-listener-go/pkg/server"
	"github.com/jc0b/fleetdm-listener-go/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var webhookCmd = &cobra.Command{
	Use:    "webhook",
	Short:  "Runs a webhook listener.",
	Long:   `Brings up a webhook listener, to receive incoming messages from webhooks.`,
	PreRun: util.PreRunSetup,
	Run: func(cmd *cobra.Command, args []string) {
		log.Trace("Trace logging will be shown...")
		log.Debug("Debug logging will be shown...")
		log.Info("Initialising webhook listener...")
		server, err := server.NewServer()
		if err != nil {
			log.Fatalf("Error initializing listener: %s", err.Error())
		}
		log.Infof("Starting server on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	},
	TraverseChildren: false,
}

func init() {
	rootCmd.AddCommand(webhookCmd)
	flags := webhookCmd.Flags()

	flags.IntP("port", "p", 8080, "The port to serve the webhook listener on.")
	viper.BindPFlag("port", flags.Lookup("port"))
}
