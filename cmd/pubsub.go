package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/pubsub"
	"github.com/jc0b/fleetdm-listener-go/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var pubsubCmd = &cobra.Command{
	Use:    "pubsub",
	Short:  "Runs a pubsub listener.",
	Long:   `Brings up a pubsub listener, to read in messages from a subscription.`,
	PreRun: util.PreRunSetup,
	Run: func(cmd *cobra.Command, args []string) {
		log.Trace("Trace logging will be shown...")
		log.Debug("Debug logging will be shown...")
		log.Info("Starting pubsub listener")

		var ctx = context.Background()
		var projectID string
		var subscriptionName string
		if viper.GetString("pubsub-subscription") != "" {
			subscriptionName = viper.GetString("pubsub-subscription")
			log.Tracef("Setting PubSub subscription name to %s due to provided configuration", subscriptionName)
		} else {
			log.Fatalf("Tried to connect to a PubSub subscription with unspecified subscription name.")
		}

		if viper.GetString("pubsub-gcp-project") != "" {
			projectID = viper.GetString("pubsub-gcp-project")
			log.Tracef("Setting PubSub project ID to %s due to provided configuration", projectID)
		} else if metadata.OnGCE() {
			log.Trace("Running on GCE... will try and detect project ID.")
			fetchedProjectID, err := metadata.ProjectIDWithContext(ctx)
			if err != nil {
				log.Fatalf("Failed to determine PubSub subscription to connect to due to an error fetching the GCP project ID from the metadata server: %s", err.Error())
			}
			projectID = fetchedProjectID
			log.Tracef("Using inferred project ID %s as the PubSub subscription project ID.", projectID)
		} else {
			log.Fatalf("Tried to connect to PubSub subscription client with empty GCP project ID.")
		}

		client, err := pubsub.NewClient(ctx, projectID)
		if err != nil {
			log.Fatal(err)
		}
		// Use a callback to receive messages via the given subscription.
		sub := client.Subscription(subscriptionName)
		fmt.Printf("Listening through subscription: %v\n", sub)
		err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
			fmt.Println(string(m.Data[:]))
			m.Ack() // Acknowledge that we've consumed the message.
		})
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Blah!")
	},
	TraverseChildren: false,
}

func init() {
	rootCmd.AddCommand(pubsubCmd)
	flags := pubsubCmd.Flags()

	flags.StringP("pubsub-subscription", "s", "", "The name of the PubSub subscription to read messages from. If unset, events will be discarded.")
	viper.BindPFlag("pubsub-subscription", flags.Lookup("pubsub-subscription"))

	flags.StringP("pubsub-gcp-project", "", "", "The name of the GCP project containing your PubSub subscription. If not set, will default to the current project you are running the listener in. If not running the listener on GCP, will default to an empty string.")
	viper.BindPFlag("pubsub-gcp-project", flags.Lookup("pubsub-gcp-project"))

}
