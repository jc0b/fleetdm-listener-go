package consumers

import (
	"context"

	"github.com/cockroachdb/errors"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/pubsub"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func RunPubSubConsumer(ctx context.Context) {
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
	defer func(client *pubsub.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)
	// Use a callback to receive messages via the given subscription.
	sub := client.Subscription(subscriptionName)
	log.Infof("Attached to subscription: %v", sub)
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		log.Infof(string(m.Data[:]))
		m.Ack() // Acknowledge that we've consumed the message.
		//TODO: do your thing here with the data
	})
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Error(err)
	}

	log.Info("Shut down PubSub consumer")
}
