package pods

import (
	"encoding/json"

	admissioncontroller "github.com/ArmedGuy/k8s-nvidia-runtime-admission-controller"

	v1 "k8s.io/api/core/v1"
)

// NewMutationHook creates a new instance of pods mutation hook
func NewMutationHook() admissioncontroller.Hook {
	return admissioncontroller.Hook{
		Create: mutateCreate(),
	}
}

func parsePod(object []byte) (*v1.Pod, error) {
	var pod v1.Pod
	if err := json.Unmarshal(object, &pod); err != nil {
		return nil, err
	}

	return &pod, nil
}
