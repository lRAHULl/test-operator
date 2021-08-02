# Build a sample Kubernetes [Operator](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/) using [kubebuilder framework](https://book.kubebuilder.io/)

1. Build Controller and deploy CRD and controller to the cluster

- `make docker-build IMG=<custom-image-name-for-controller>`
- `make docker-push IMG=<custom-image-name-for-controller>`
- `make deploy IMG=<custom-image-name-for-controller>`
