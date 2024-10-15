# How to create a kubebuilder project

# Initialize kubebuilder project

`kubebuilder init --domain <your domain> --repo <your domain>/<module-name>`

ex:

`kubebuilder init --domain hemmelig.io --repo github.com/rogerwesterbo/k8s-notifier`

## Add a kubebuilder api with custom resource definition (CRD)

`kubebuilder create api --group task --version v1 --kind Notify`

## Adjust types

Go to file `src/api/v1/notify_types.go` and add more properties to you needs

Example:

```golang
type NotifySpec struct {
	Name string          `json:"name,omitempty"`
	Type metav1.TypeMeta `json:"type,omitempty"`
}

type NotifyStatus struct {
	Notified bool `json:"notified"`
    NotifiedTimestamp string `json:"notifiedTomestamp"`
}

type Notify struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NotifySpec   `json:"spec,omitempty"`
	Status NotifyStatus `json:"status,omitempty"`
}
```

## Build

`make`

## Create CRD file

`make manifests`

## Create a cluster or reuse an existing

If you already have a cluster, logon to that cluster and check if you see noeds:
ex:

```bash
kubectl get nodes
NAME                        STATUS   ROLES           AGE   VERSION
testcluster-control-plane   Ready    control-plane   61s   v1.31.0
testcluster-worker          Ready    <none>          45s   v1.31.0
testcluster-worker2         Ready    <none>          45s   v1.31.0
```

Or create a kind cluster locally by [reading this](./k8s-locally.md)

## install CRD to cluster

`make install`
Installs the CRD and can be viewed with kubectl like this:

```bqsh
kubectl api-resources
...
notifies                                         task.hemmelig.io/v1               true         Notify
...
```

To see more documentation with kubectl:

```bash
kubectl explain notifies
GROUP:      task.hemmelig.io
KIND:       Notify
VERSION:    v1

DESCRIPTION:
    Notify is the Schema for the notifies API

FIELDS:
  apiVersion    <string>
    APIVersion defines the versioned schema of this representation of an object.
    Servers should convert recognized schemas to the latest internal value, and
    may reject unrecognized values. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources

  kind  <string>
    Kind is a string value representing the REST resource this object
    represents. Servers may infer this from the endpoint the client submits
    requests to. Cannot be updated. In CamelCase. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds

  metadata      <ObjectMeta>
    Standard object's metadata. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata

  spec  <Object>
    NotifySpec defines the desired state of Notify

  status        <Object>
    NotifyStatus defines the observed state of Notify
```

Or to see more under notifies:

```bash
kubectl explain notifies.status
GROUP:      task.hemmelig.io
KIND:       Notify
VERSION:    v1

FIELD: status <Object>

DESCRIPTION:
    NotifyStatus defines the observed state of Notify

FIELDS:
  notified      <boolean>
    <no description>

  notifiedTimestamp     <string>
    <no description>
```

## Add a CRD

```bash
kubectl apply -f manifests/notify-test.yaml

notify.task.hemmelig.io/test-notify created
```

```bash
kubectl get notifies -A

NAMESPACE   NAME          AGE
default     test-notify   23s
```
