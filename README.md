# k8s object churner (koc)

A utility project for simulating object creation churn on a kubernetes cluster.

## Table of Contents

- [Running locally](#running-locally)
- [Usage](#usage)
- [Support](#support)
- [Contributing](#contributing)

## Running-locally

- On Docker:
    - Needs go 1.17 or above setup
    - Building ` make build`
    - Running `make run`
- On Mac:
    - Needs Docker runtime
    - Building ` make docker`
    - Running `make docker-run`
- On k8s:
    - Needs Deployment access to k8s cluster
    - Set cluster context using `kubeconfig config use-context <context>`
    - Building ` make k8s-install`
    - Running `make k8s-uninstall`

## Usage

## Support

Please [open an issue](https://github.com/fraction/readme-boilerplate/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits,
and [open a pull request](https://github.com/fraction/readme-boilerplate/compare/).
