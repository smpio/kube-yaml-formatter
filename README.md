# kube-yaml-formatter

Reformats YAML as Kubernetes does


## Installation

Clone the repo and run `make install` to install it to `/usr/local/bin`. Docker is required to build.


## Details

Kubernetes itself [uses](https://github.com/kubernetes/kubernetes/blob/528859ef9df0434f8dedc337e24628427f480b17/staging/src/k8s.io/cli-runtime/pkg/printers/yaml.go) [sigs.k8s.io/yaml](https://sigs.k8s.io/yaml), which is wrapper for [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2). `sigs.k8s.io/yaml` doesn't support saving map keys order, but `gopkg.in/yaml.v2` do. So we use `gopkg.in/yaml.v2` here.
