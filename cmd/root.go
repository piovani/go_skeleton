package cmd

import (
	"fmt"
	"log"

	"github.com/piovani/go_skeleton/infra/config"
	"github.com/spf13/cobra"
)

func Execute() {
	cmd := &cobra.Command{
		Use:     "go_skeleton",
		Version: "1.0.0",
	}

	InitConfig()

	cmd.AddCommand(
		API,
	)

	CheckFatal(cmd.Execute())
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(fmt.Printf("error during application startup: %v", err))
	}
}

func InitConfig() {
	CheckFatal(config.InitConfig())
}
