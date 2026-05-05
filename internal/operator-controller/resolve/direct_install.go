package resolve

import (
	"context"

	ocv1 "github.com/operator-framework/operator-controller/api/v1"
	"github.com/operator-framework/operator-controller/internal/operator-controller/bundle"
	"github.com/operator-framework/operator-registry/alpha/declcfg"
)

type DirectResolver struct{}

func (r *DirectResolver) Resolve(ctx context.Context, ext *ocv1.ClusterExtension, installedBundle *ocv1.BundleMetadata) (*declcfg.Bundle, *bundle.VersionRelease, *declcfg.Deprecation, error) {
	return &declcfg.Bundle{
			Name:    "direct-install",
			Package: "direct-install",
			Image:   ext.Spec.Source.OCIImage.Ref,
		},
		&bundle.VersionRelease{},
		&declcfg.Deprecation{},
		nil
}
