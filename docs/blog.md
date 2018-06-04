# My first Kubernetes based platform

Goal: take an existing solution and convert it to a Kubernetes aware platform / SaaS. Do this with minimal modification to existing code. Any modification to the code should not require that the software run in Kubernetes.

## Requirements

- Software should be fully deployed and configured using Helm.
- Due to requirement 1, all Azure services should be deployed using the Kubernetes Service Catalog and an Azure broker.
- Deployment, management, and observation of many concurrent instances should be simple without scripting, complex queries.
- The solution should have a REST interfact for deploy, delete, list, and show functions.
- Ultimately, Kubernetes should be obfuscated from the solution.

## Kubernetes Service Catalog / Azure Open Service Broker

Because the solution uses three Azure services (Storage Queue, CosmosDB, and text sentiment API), these services need to be deployed and consumable using the Kubernetes Service Catalog. This is necessary so that a Helm chart can be written for complete solution deployment.

**Lessons Learned:**

- Limited capability in some current OSBA modules. Specifically for this project the OSBA storage module cannot create a Storage Queue. This was easy to work around in the application code as can be seen in [this commit](https://github.com/neilpeterson/twitter-sentiment-for-kubernetes/commit/a14024bae39e17e2ea198fe43d20a02b5c308ab9#diff-2b67715d77acf2f5347a37dfc08c362a).
- Limited available OSBA modules. Specifically, this project requires an Azure text sentiment analysis API. OSBA does not have a module for Cognitive Services. This module was relatively simple to create and now exists in this [OSBA fork](https://github.com/neilpeterson/open-service-broker-azure/tree/cognitive-services-text-analytics).

## Helm Chart

Helm to install and uninstall instances of the Twitter Sentiment SaaS. I like that this allows the logic of the software to be removed from the operator itself.

**Lessons Learned:**

- I'm not 100% on how to best run Helm releases.via Go code. In the interim, the solution uses a gross hack that involves starting a Kubernetes job, which runs Helm commands. I will update this routine as time allows.
- I have tried a few methods for tracking / observing individual instances of the tweet-factory service. The current solution is to use randomly generated Helm release names, and add this name as a label to each object.

## Operator

I've used the CoreOS / Redhat operator-framework for creating the tweet-factory operator.

## What's next

I'm currently working on a REST interface / CLI for starting, deleting, and observing instances of the tweet-factory service.