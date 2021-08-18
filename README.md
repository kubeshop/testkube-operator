# TOC 

- [KubeTest-operator - your testing friend](#kubetest-operator---your-testing-friend)

# KubeTest-operator - your testing friend

This is the operator for Kubernetes-native framework for definition and execution of tests in a cluster; 

*`Please note!`* For now it has limited functionality and only installs needed CRDs (custom resurces definitions) into an active k8s cluster. It's meant to be installed as a part of main chart from here: https://github.com/kubeshop/kubetest/

Instead of orchestrating and executing test with a CI tool (jenkins, travis, circle-ci, GitHub/GitLab, etc) tests are defined/orchestrated in the cluster using k8s native concepts (manifests, etc) and executed automatically when target resources are updated in the cluster. Results are written to existing tooling (prometheus, etc). This decouples test-definition and execution from CI-tooling/pipelines and ensures that tests are run when corresponding resources are updated (which could still be part of a CI/CD workflow). 