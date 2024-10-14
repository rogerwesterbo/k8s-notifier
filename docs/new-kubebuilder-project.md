# How to create a kubebuilder project

# Initialize kubebuilder project

`kubebuilder init --domain <your domain> --repo <your domain>/<module-name>`

ex:

`kubebuilder init --domain hemmelig.io --repo github.com/rogerwesterbo/k8s-notifier`

## Add a kubebuilder api with custom resource definition (CRD)

`kubebuilder create api --group task --version v1 --kind Notify`

## Build

`make`

## Create CRD file

`make manifests`
