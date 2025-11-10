# FleetDM event listener
A proof-of-concept listener for FleetDM events. Can operate in two modes:
- PubSub listener that reads from a specified topic, and can operate on the messages
- Webhook listener that listens on a specified port, and can operate on the consumed webhooks

You could theoretically run this as a pod within a Kubernetes deployment, so all of your FleetDM deployment stuff stays within a given cluster.


I'd love to evolve this in the future to be one binary that can do all of these things at once, but at the moment you'll need an instance for webhooks, and an instance per-PubSub subscription.