# Kubernetes Nvidia Runtime Admission Controller
Based on the boilerplate [admission controller by douglasmakey](https://github.com/douglasmakey/admissioncontroller).

This controller has one simple purpose: set runtimeClassName: nvidia for any pod that is scheduled with Nvidia GPU resources.

This is especially useful for when systems like Kubeflow is deployed to heterogenous clusters where the nvidia runtime cannot be the default across all nodes, or for installations like k3s where nvidia runtime is registered as a separate runtime class.


## Deployment
The project currently supports deployment to existing kubeflow projects that use cert-manager, by using kustomize.

`kustomize build kustomize/overlays/cert-manager | kubectl apply -f` is the easiest way to install the admission controller to the `default` namespace.


## Verify it works
In a cluster that contains a runtime class for `nvidia` (such as k3s installed on systems with GPU, or other custom installations, probably not using gpu-operator that already does this), deploy a pod asking for nvidia gpus:
```
$ cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: gpu-pod
spec:
  restartPolicy: Never
  containers:
    - name: cuda-container
      image: nvcr.io/nvidia/k8s/cuda-sample:vectoradd-cuda10.2
      resources:
        limits:
          nvidia.com/gpu: 1 # requesting 1 GPU
  tolerations:
  - key: nvidia.com/gpu
    operator: Exists
    effect: NoSchedule
EOF
```

As a fun side-effect, by letting this webhook set the runtime without the user specifying it, the warning in https://github.com/NVIDIA/k8s-device-plugin regarding all GPUs being exposed if you don't explicitly request resources will not be a problem as containers won't even use the nvidia runtime unless they do request resources.