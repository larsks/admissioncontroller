package annotations

import (
	"github.com/douglasmakey/admissioncontroller"

	"k8s.io/api/admission/v1beta1"
	log "k8s.io/klog/v2"
)

func validateCreate() admissioncontroller.AdmitFunc {
	return func(r *v1beta1.AdmissionRequest) (*admissioncontroller.Result, error) {
		thing, err := parseAnything(r.Object.Raw)
		if err != nil {
			return &admissioncontroller.Result{
				Msg: err.Error(),
			}, nil
		}

		log.Infof("thing: %+v", thing)
		log.Infof("request: %+v", r)

		return &admissioncontroller.Result{Allowed: true}, nil
	}
}
