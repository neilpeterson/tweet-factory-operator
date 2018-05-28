# Tweet Factory - A sample SaaS built on Kubernetes

Tweet Factory is a sample / reference Kubernetes operator that demonstrates running [this Twitter sentiment analysis solution](https://github.com/neilpeterson/twitter-sentiment-for-kubernetes) as a service on top of Kubernetes.

I've built this project to demonstrate concepts for building platforms on Kubernetes. The following technology is used in this solution:

- Kubernetes Custom Resources (TweetFactory)
- Kubernetes Service Catalog
- Open Service Broker for Azure
- Operator Framework

## Demonstration

Short five minute demo – this is also available on [youtube.com](https://youtu.be/_WWMCsYPvBQ), which may be easier to view from the beginning, pause, etc.

![](images/tweet-factory.gif)

## Prerequisites

In order to run Tweet Factory, you need:

- Kubernete cluster with Helm configured
- Service Catalog installed on the cluster
- Open Service Broker for Azure (customized with [this fork](https://github.com/neilpeterson/open-service-broker-azure-samples/tree/master/osba-text-analytics))
- Twitter API Consumer key and access token

## Installation

Add this Helm repo:

```
helm repo add azure-samples https://azure-samples.github.io/helm-charts/
```

Start the operator:

```
helm install azure-samples/tweet-factory-operator -n tweet-factory
```

## Run the solution

At this point, a custom resource definition (CRD) has been created that represents the Tweet Factory solution. When a new instance of the custom resource is created, the Tweet Factor operator creates an instance of the service.

To deploy a new instance, create a file named `tweet-seattle.yaml` and copy in the following YAML. Update the Twitter consumer key, access token, and both secrets with the values from your twitter API application.

```
apiVersion: "tweet-factory.com/v1alpha1"
kind: "TweetFactory"
metadata:
  name: "seattle"
spec:
  consumerKey: ""
  consumerSecret: ""
  accessToken: ""
  accessTokenSecret: ""
  filterText: "Seattle"
```

Create the custom resource with `kubectl apply`.

```
kubectl apply -f tweet-seattle.yaml
```

Once the resource is created, the Tweet Sentiment operator runs a Helm chart the following items are created:

- Azure Storage Queue - tweets to be processed are stored here
- Azure Cosmos DB - sentiment results are stored here
- Azure Text Analytics API - performs sentiment analysis
- Get tweet POD - gets tweets and stores them in Queue
- Process tweet POD - performs sentiment analysis and stores result in Cosmos DB
- Chart tweet POD - visualizes the results