package resolve

import (
	"context"
	"fmt"

	ocv1 "github.com/operator-framework/operator-controller/api/v1"
	"github.com/operator-framework/operator-controller/internal/operator-controller/bundle"
	"github.com/operator-framework/operator-registry/alpha/declcfg"
)

type UnionResolver struct {
	resolver map[string]Resolver
}

func (r *UnionResolver) Resolve(ctx context.Context, ext *ocv1.ClusterExtension, installedBundle *ocv1.BundleMetadata) (*declcfg.Bundle, *bundle.VersionRelease, *declcfg.Deprecation, error) {
	res, ok := r.resolver[ext.Spec.Source.SourceType]
	if !ok || res == nil {
		return nil, nil, nil, fmt.Errorf("unrecognized source type %s for clusterextension %s", ext.Spec.Source.SourceType, ext.Name)
	}
	return res.Resolve(ctx, ext, installedBundle)
}

func (r *UnionResolver) Register(key string, resolver Resolver) {
	if r.resolver == nil {
		r.resolver = map[string]Resolver{}
	}
	r.resolver[key] = resolver
}
