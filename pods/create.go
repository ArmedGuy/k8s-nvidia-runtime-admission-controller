package pods

import (
	admissioncontroller "github.com/ArmedGuy/k8s-nvidia-runtime-admission-controller"

	"k8s.io/api/admission/v1beta1"
)

func mutateCreate() admissioncontroller.AdmitFunc {
	return func(r *v1beta1.AdmissionRequest) (*admissioncontroller.Result, error) {
		var operations []admissioncontroller.PatchOperation
		pod, err := parsePod(r.Object.Raw)
		if err != nil {
			return &admissioncontroller.Result{Msg: err.Error()}, nil
		}

		gpuRequested := false
		for _, container := range pod.Spec.Containers {
			for limit, _ := range container.Resources.Limits {
				if limit == "nvidia.com/gpu" {
					gpuRequested = true
					break
				}
			}
		}
		if gpuRequested {
			operations = append(operations, admissioncontroller.ReplacePatchOperation("/spec/runtimeClassName", "nvidia"))
		}

		return &admissioncontroller.Result{
			Allowed:  true,
			PatchOps: operations,
		}, nil
	}
}
