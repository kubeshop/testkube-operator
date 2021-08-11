# TOC 

- [KubeTest-operator - your testing friend](#kubetest-operator---your-testing-friend)
- [Helm installation](#helm-installation)
    * [Helm defaults used with the chart](#helm-defaults-used-with-the-chart)

# KubeTest-operator - your testing friend

This is the operator for Kubernetes-native framework for definition and execution of tests in a cluster; 

*`Please note!`* For now it has limited functionality and only installs needed CRDs (custom resurces definitions) into an active k8s cluster. It's meant to be installed as a part of main chart from here: https://github.com/kubeshop/kubetest/

Instead of orchestrating and executing test with a CI tool (jenkins, travis, circle-ci, GitHub/GitLab, etc) tests are defined/orchestrated in the cluster using k8s native concepts (manifests, etc) and executed automatically when target resources are updated in the cluster. Results are written to existing tooling (prometheus, etc). This decouples test-definition and execution from CI-tooling/pipelines and ensures that tests are run when corresponding resources are updated (which could still be part of a CI/CD workflow). 

# Helm installation

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:
```sh
helm repo add kubetest-operator https://kubeshop.github.io/kubetest-operator
```

If you had already added this repo earlier, run `helm repo update` to retrieve
the latest versions of the packages.  You can then run `helm search repo
kubetest-operator` to see the charts.

To install the kubetest-operator chart:
```sh
helm install my-<chart-name> kubetest-operator/kubetest-operator
```
To uninstall the kubetest-operator chart:
```sh
helm delete my-<chart-name> kubetest-operator/kubetest-operator
```
> Please note that this Helm chart will install only CRDs and nothing else for now. It's meant to be installed as a part of main chart from here: https://github.com/kubeshop/kubetest/