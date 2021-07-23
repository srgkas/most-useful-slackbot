### Deploy
k8s cluster should been created already.
For local environment `minikube start` can be used.

```
helm repo add bitnami https://charts.bitnami.com/bitnami > /dev/null
helm install --set architecture=standalone,auth.enabled=false redis bitnami/redis
helm install --set slack_token=SOME_TOKEN --set github_token=ANOTHER_TOKEN --set redis.host=redis-master slack-bot charts/app/
```

Delete pods & services than `helm` has created:
```
helm delete slack-bot && helm delete redis
```
