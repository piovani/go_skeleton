package cmd

import (
	"github.com/piovani/go_skeleton/api"
	"github.com/spf13/cobra"
)

var (
	API = &cobra.Command{
		Use:     "api",
		Short:   "Start Api Resquet",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			CheckFatal(api.NewApi().Start())
		},
	}
)
