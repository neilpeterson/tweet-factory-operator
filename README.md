# Twitter Sentiment Operator - a sample platform built on Kubernetes

## Situation

Your orginization has developed a framework for analysing, storing, and charting the sentiment of Tweets based on a key word. Your orginzation woudl like to provide the software as an on demand serivce for users both inside and outside of the orginzation. The following requiremets shoudl be met:

- Twitter sentiment framework should run on Kuberntes
- Consumers should not need any Kuberntes experiance
- Multiple instances may be runnign at any given time

## My observations

- Application should be fully baked into a Helm chart
- Service catalog integration for mnaged services
