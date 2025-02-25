package features

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/tabwriter"

	"k8s.io/apimachinery/pkg/util/json"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/featuregate"
	"sigs.k8s.io/yaml"
)

const (
	// Add new feature gates constants (strings)
	// Ex: SomeFeature featuregate.Feature = "SomeFeature"
	PreflightPermissions featuregate.Feature = "PreflightPermissions"
)

type FeatureGateInfo struct {
	featuregate.FeatureSpec
	name        featuregate.Feature
	description string
}

var operatorControllerFeatureGates = []FeatureGateInfo{
	// Add new feature gate definitions
	// Ex: SomeFeature: {...}
	{
		FeatureSpec: featuregate.FeatureSpec{
			Default:       false,
			PreRelease:    featuregate.Alpha,
			LockToDefault: false,
		},
		name:        PreflightPermissions,
		description: "Enable checking permissions preflight",
	},
}

var OperatorControllerFeatureGate featuregate.MutableFeatureGate = featuregate.NewFeatureGate()

func init() {
	sort.Slice(operatorControllerFeatureGates, func(i, j int) bool {
		return operatorControllerFeatureGates[i].name < operatorControllerFeatureGates[j].name
	})
	featureSpecs := map[featuregate.Feature]featuregate.FeatureSpec{}
	for _, featureInfo := range operatorControllerFeatureGates {
		featureSpecs[featureInfo.name] = featureInfo.FeatureSpec
	}
	utilruntime.Must(OperatorControllerFeatureGate.Add(featureSpecs))
}

func PrintFeatureGateHelp(format string) (string, error) {
	type featureHelp struct {
		Name        featuregate.Feature
		Stability   string
		Enabled     bool
		Description string
	}
	var catalogdFeatureHelp []featureHelp
	for _, featureInfo := range operatorControllerFeatureGates {
		catalogdFeatureHelp = append(catalogdFeatureHelp, featureHelp{
			Name:        featureInfo.name,
			Stability:   string(featureInfo.PreRelease),
			Enabled:     featureInfo.Default, // TODO: special handling for deprecated features?
			Description: featureInfo.description,
		})
	}

	switch strings.ToLower(format) {
	case "json":
		out, err := json.Marshal(catalogdFeatureHelp)
		return string(out), err
	case "yaml":
		out, err := yaml.Marshal(catalogdFeatureHelp)
		return string(out), err
	default:
		out := &bytes.Buffer{}
		tw := tabwriter.NewWriter(out, 8, 4, 4, ' ', 0)
		fmt.Fprintf(tw, "Name\tStability level\tEnabled by Default\tDescription\n")
		for _, featureInfo := range catalogdFeatureHelp {
			fmt.Fprintf(tw, "%s\t%s\t%v\t%s\n", featureInfo.Name, featureInfo.Stability, featureInfo.Enabled, featureInfo.Description)
		}
		err := tw.Flush()
		return out.String(), err
	}
}
