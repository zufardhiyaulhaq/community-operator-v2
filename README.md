# community-operator-v2

Manage Community lifecycle via Kubernetes CRDs

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square) [![made with Go](https://img.shields.io/badge/made%20with-Go-brightgreen)](http://golang.org) [![Github master branch build](https://img.shields.io/github/workflow/status/zufardhiyaulhaq/community-operator-v2/Master)](https://github.com/zufardhiyaulhaq/community-operator-v2/actions/workflows/master.yml) [![GitHub issues](https://img.shields.io/github/issues/zufardhiyaulhaq/community-operator-v2)](https://github.com/zufardhiyaulhaq/community-operator-v2/issues) [![GitHub pull requests](https://img.shields.io/github/issues-pr/zufardhiyaulhaq/community-operator-v2)](https://github.com/zufardhiyaulhaq/community-operator-v2/pulls)[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/community-operator-v2)](https://artifacthub.io/packages/search?repo=community-operator-v2)

## Concept
community-operator-v2 is a Kubernetes operator. This operator watch and reconcile the custom resources to manage the community lifecycle.

currently supported:
- [x] Support Telegram Handler
- [x] Support Twitter Handler
- [ ] Support Slack Handler
- [ ] Support Discord Handler
- [ ] Support Facebook Handler
- [ ] Support Instagram Handler
- [x] Manage and broadcast community message like weekly, announcement, and meetup automatically.
- [ ] Auto generate Meetup poster
- [ ] Auto generate community website and populate community message like weekly, announcement and meetup.

#### Community CRD
Define a community, you can create multiple community object to manage multiple community.
```
apiVersion: community.zufardhiyaulhaq.com/v1alpha1
kind: Community
metadata:
  name: community-sample
spec:
  socialMedia:
    telegram:
    - telegramhandler-sample
    twitter:
    - twitterhandler-sample
```
You can define list of social media that this community run. Any message like weekly, announcement, and meetup will automatically broadcasted in all the community social media. `spec.socialMedia.telegram` is a list of TelegramHandler object and `spec.socialMedia.twitter` is a list of TwitterHandler object.

#### TelegramHandler CRD
Define an entrypoint to access the community telegram channel or group.
```
apiVersion: community.zufardhiyaulhaq.com/v1alpha1
kind: TelegramHandler
metadata:
  name: telegramhandler-sample
spec:
  type: group
  credential: "-9999999999"
  authentication:
    token:
      secret:
        name: telegramhandler-secret-sample
        key: token
```
`spec.type` support both `group` and `channel`.

#### TwitterHandler CRD
Define an entrypoint to access the community twitter account.
```
apiVersion: community.zufardhiyaulhaq.com/v1alpha1
kind: TwitterHandler
metadata:
  name: twitterhandler-sample
spec:
  authentication:
    apiKey:
      secret:
        name: twitterhandler-secret-sample
        key: apiKey
    apiKeySecret:
      secret:
        name: twitterhandler-secret-sample
        key: apiKeySecret
    accessToken:
      secret:
        name: twitterhandler-secret-sample
        key: accessToken
    accessTokenSecret:
      secret:
        name: twitterhandler-secret-sample
        key: accessTokenSecret
```

#### Announcement CRD
You can send custom announcement to your community via Announcement CRD.
```
apiVersion: community.zufardhiyaulhaq.com/v1alpha1
kind: Announcement
metadata:
  name: announcement-sample
spec:
  community:
  - "community-sample"
  spec:
    subject: Your subject here
    body: >-
      Your Announcement Body here
    imageUrl: https://your.image.here
    tags:
      - list
      - of
      - tag
```
You can select multiple Community object under `spec.community`.

#### Meetup CRD
You can send Meetup message to your community.
```
apiVersion: community.zufardhiyaulhaq.com/v1alpha1
kind: Meetup
metadata:
  name: meetup-sample
spec:
  community:
  - "community-sample"
  spec:
    name: "name"
    date: "date"
    time: "time"
    place: "place"
    registrationUrl: https://your.registration.url.here
    imageUrl: https://your.image.here
    tags:
      - list
      - of
      - tag
    sponsors:
    - name: sponsor-name
    speakers:
      - name: speaker-name
        position: speaker-position
        company: speaker-company
        title: talk-title
```
You can select multiple Community object under `spec.community`.

#### Weekly CRD
Weekly is a update about your community. In tech community, things keep changing every week and we need a specific handler to send weekly update to the community.
```
apiVersion: community.zufardhiyaulhaq.com/v1alpha1
kind: Weekly
metadata:
  name: weekly-sample
spec:
  community:
  - "community-sample"
  spec:
    name: "name"
    date: "date"
    imageUrl: https://your.image.here
    tags:
      - list
      - of
      - tag
    articles:
    - title: article-title
      url: https://your.article.url.here
      type: article-type
```
You can select multiple Community object under `spec.community`.

## Installing

To install the chart with the release name `my-release`:

```console
helm repo add community-operator-v2 https://zufardhiyaulhaq.com/community-operator-v2/charts/releases/
helm install my-community-operator-v2 community-operator-v2/community-operator-v2 --values values.yaml
```

## Usage
1. Apply some example
```console
kubectl apply -f examples/
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| operator.image | string | `"zufardhiyaulhaq/community-operator-v2"` |  |
| operator.replica | int | `1` |  |
| operator.tag | string | `"v1.0.0"` |  |
| resources.limits.cpu | string | `"200m"` |  |
| resources.limits.memory | string | `"100Mi"` |  |
| resources.requests.cpu | string | `"100m"` |  |
| resources.requests.memory | string | `"20Mi"` |  |

see example files [here](https://github.com/zufardhiyaulhaq/community-operator-v2/blob/master/charts/community-operator-v2/values.yaml)

