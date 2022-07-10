# community-operator-v2

Expose your service in Kubernetes to the Internet with open source FRP!

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square) [![made with Go](https://img.shields.io/badge/made%20with-Go-brightgreen)](http://golang.org) [![Github master branch build](https://img.shields.io/github/workflow/status/zufardhiyaulhaq/community-operator-v2/Master)](https://github.com/zufardhiyaulhaq/community-operator-v2/actions/workflows/master.yml) [![GitHub issues](https://img.shields.io/github/issues/zufardhiyaulhaq/community-operator-v2)](https://github.com/zufardhiyaulhaq/community-operator-v2/issues) [![GitHub pull requests](https://img.shields.io/github/issues-pr/zufardhiyaulhaq/community-operator-v2)](https://github.com/zufardhiyaulhaq/community-operator-v2/pulls)[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/community-operator-v2)](https://artifacthub.io/packages/search?repo=community-operator-v2)

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

