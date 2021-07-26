## most-useful-slackbot

### Requirements
This bot *currently* exposes Github and Slack APIs to automate some routine workflows.

1. Github
In order use Github API the token should be obtained.
Visit (https://github.com/settings/tokens)[https://github.com/settings/tokens]
and set it to env variable: `GITHUB_TOKEN`. See `.env.example` for more info.

2. Slack
Yes :) Slack bot should created via (https://api.slack.com/apps)[https://api.slack.com/apps] and installed to organization.
Set obtained token to `SLACK_TOKEN` env var. See `.env.example` for more info.

Following OAuth scopes are **required**:
```
- channels:history
- channels:join
- channels:manage
- chat:write
- emoji:read
- im:history
- incoming-webhook
- reactions:read
- channels:read
- groups:read
- im:read
- mpim:read
- commands
```

This bot is built on top of next Slack API features:
- Event subscriptions
- Interactivity
- Slash commands

**Event subscriptions**
Request url MUST be specified. 
E.g.: `https://mydomain.local/events/handle`

Bot reacts on following events:
```
- emoji_changed
- message.channels
- message.im
- reaction_added
- reaction_removed
```

**Interactivity**
Request url MUST be specified. 
E.g.: `https://mydomain.local/interact`

**Slash commannds**
Command name (e.g `/hotfix`) and request url MUST be specified.
E.g.: `https://mydomain.local/commands/hotfix`


### Development
Development process requires several tools:
- `go 1.16`
- `docker`
- `docker-compose`
- `ngrok`
- `minikube`
- `helm`

### Configuration
Sanity check:
- Make sure api tokens `GITHUB_TOKEN` and `SLACK_TOKEN` are present and valid
- Check that channels from config are created and bot has been added there. See `.env.example`
for channels configuration.

### Features
- TBD


### Deployment
Kubernetes cluster should be created already.

For local environment `minikube start` can be used.

```
helm repo add bitnami https://charts.bitnami.com/bitnami > /dev/null
helm install --set architecture=standalone,auth.enabled=false redis bitnami/redis
helm install --set slack_token=SOME_TOKEN --set github_token=ANOTHER_TOKEN --set redis.host=redis-master slack-bot charts/app/
```

Delete pods & services that `helm` has created:
```
helm delete slack-bot && helm delete redis
```