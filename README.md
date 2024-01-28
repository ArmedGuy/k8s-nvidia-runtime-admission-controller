# Kubernetes Nvidia Runtime Admission Controller
Based on the boilerplate [admission controller by douglasmakey](https://github.com/douglasmakey/admissioncontroller).

This controller has one simple purpose: set runtimeClassName: nvidia for any pod that is scheduled with Nvidia GPU resources.

This is especially useful for when systems like Kubeflow is deployed to heterogenous clusters where nvidia runtime cannot be the default across all nodes.


## Deployment
The project currently supports deployment to existing kubeflow projects that use cert-manager, by using kustomize.

`kustomize build kustomize/overlays/cert-manager | kubectl apply -f` is the easiest way to install the admission controller to the `default` namespace.