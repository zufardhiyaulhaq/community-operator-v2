# Contributing
By participating to this project, you agree to abide our [code of conduct](https://github.com/zufardhiyaulhaq/community-operator-v2/blob/master/.github/CODE_OF_CONDUCT.md).

## Development
For small things like fixing typos in documentation, you can [make edits through GitHub](https://help.github.com/articles/editing-files-in-another-user-s-repository/), which will handle forking and making a pull request (PR) for you. For anything bigger or more complex, you'll probably want to set up a development environment on your machine, a quick procedure for which is as folows:

### Setup your machine
Prerequisites:
- make
- [Go 1.17](https://golang.org/doc/install)
- [operator-sdk v1.21.0](https://sdk.operatorframework.io/)

Fork and clone **[community-operator-v2](https://github.com/zufardhiyaulhaq/community-operator-v2)** repository.

- deploy CRDs
```
kubectl apply -f config/crd/bases/
```

- Run community-operator-v2 locally
```
make install run
```

- deploy some examples
```
kubectl apply -f examples/deployment/
kubectl apply -f examples/client/
```

### Submit a pull request
As you are ready with your code contribution, push your branch to your `community-operator-v2` fork and open a pull request against the **master** branch.

Please also update the [CHANGELOG.md](https://github.com/zufardhiyaulhaq/community-operator-v2/blob/master/CHANGELOG.md) to note what you've added or fixed.
