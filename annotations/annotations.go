package annotations

import (
	"encoding/json"

	"github.com/douglasmakey/admissioncontroller"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewValidationHook creates a new instance of pods validation hook
func NewValidationHook() admissioncontroller.Hook {
	return admissioncontroller.Hook{
		Create: validateCreate(),
	}
}

func parseAnything(object []byte) (*metav1.PartialObjectMetadata, error) {
	var thing metav1.PartialObjectMetadata
	if err := json.Unmarshal(object, &thing); err != nil {
		return nil, err
	}

	return &thing, nil
}
