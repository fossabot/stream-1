package kubeprometheus

import "github.com/merico-dev/stream/internal/pkg/util/helm"

// validate validates the options provided by the core.
func validate(param *helm.HelmParam) []error {
	return helm.Validate(param)
}
