/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/nqnlong1506/go_toolkit/internal"
	"github.com/spf13/cobra"
)

var serviceNames []string

// microservicesCmd represents the microservices command
var microservicesCmd = &cobra.Command{
	Use:   "microservices",
	Short: "Init a  Microservices Application",
	Long: `Initialize (go_toolkit microservices) will init folder-structure for a microservices application.

microservices must be run inside of a go module (please run "go mod init <MODNAME>" first)`,
	Run: func(cmd *cobra.Command, args []string) {
		// go mod init
		err := internal.MicroservicesInit()
		if err != nil {
			fmt.Println("Error initializing microservices:", err)
			return
		}

		// generate folder structure
		err = internal.MicroservicesGenerateFolderStructure(serviceNames)
		if err != nil {
			fmt.Println("Error generating microservices folder structure:", err)
			return
		}

		// generate Dockerfiles
		internal.MicroservicesGenerateDockerfiles(serviceNames)

	},
}

func init() {
	rootCmd.AddCommand(microservicesCmd)

	microservicesCmd.Flags().StringArrayVarP(&serviceNames, "service", "s", nil, "Name of the service to initialize")
}
