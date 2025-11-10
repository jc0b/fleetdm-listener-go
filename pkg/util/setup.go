package util

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func PreRunSetup(cmd *cobra.Command, args []string) {
	if viper.GetBool("json-logging") {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}
	if viper.GetBool("trace") {
		log.SetLevel(log.TraceLevel)
	} else if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
