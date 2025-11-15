# FleetDM event listener
A proof-of-concept listener for FleetDM events. Has two components:
- PubSub listener that reads from a specified topic, and can operate on the messages
- Webhook listener that listens on a specified port, and can operate on the consumed webhooks

You could theoretically run this as a pod within a Kubernetes deployment, so all of your FleetDM deployment stuff stays within a given cluster.


## Docs
```
FleetDM Listener is an example of a multi-use binary for tracking FleetDM events, via PubSub or webhooks.

Usage:
  fleetdm-listener [flags]

Flags:
      --debug                        Sets log level to debug. If multiple log level flags are set, the most verbose option will be respected.
  -h, --help                         help for fleetdm-listener
      --json-logging                 Enables JSON logging when set.
  -p, --port int                     The port to serve the webhook listener on. (default 8080)
      --pubsub-gcp-project string    The name of the GCP project containing your PubSub subscription. If not set, will default to the current project you are running the listener in. If not running the listener on GCP, will default to an empty string.
  -s, --pubsub-subscription string   The name of the PubSub subscription to read messages from. If unset, events will be discarded.
      --trace                        Sets log level to trace. If multiple log level flags are set, the most verbose option will be respected
```