package main

import (
	"fmt"

	"github.com/operator-framework/operator-controller/internal/catalogd/features"
	"github.com/spf13/cobra"
)

var (
	outputFormat string
)

func addFeatureGateHelpCmd(cmd *cobra.Command) {
	featureGateHelpCommand := &cobra.Command{
		Use:   "list-feature-gates",
		Short: "List information about known feature gates",
		RunE: func(_ *cobra.Command, _ []string) error {
			out, err := features.PrintFeatureGateHelp(outputFormat)
			fmt.Println(out)
			return err
		},
	}
	featureGateHelpCommand.Flags().StringVar(&outputFormat, "output", "tsv", "output format. Supported formats: 'json', 'yaml', 'tsv'")
	cmd.AddCommand(featureGateHelpCommand)
}
