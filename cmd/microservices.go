/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var serviceName []string

// microservicesCmd represents the microservices command
var microservicesCmd = &cobra.Command{
	Use:   "microservices",
	Short: "Init a  Microservices Application",
	Long: `Initialize (go_toolkit microservices) will init folder-structure for a microservices application.

microservices must be run inside of a go module (please run "go mod init <MODNAME>" first)`,
	Run: func(cmd *cobra.Command, args []string) {
		generatePaths := func(service string) []string {
			return []string{
				fmt.Sprintf("proto/%s", service),
				fmt.Sprintf("services/%s/cmd", service),
				fmt.Sprintf("services/%s/internal", service),
				fmt.Sprintf("services/%s/internal/app", service),
				fmt.Sprintf("services/%s/internal/domain", service),
				fmt.Sprintf("services/%s/internal/infrastructure", service),
				fmt.Sprintf("services/%s/internal/infrastructure", service),
				fmt.Sprintf("services/%s/internal/infrastructure/repository", service),
				fmt.Sprintf("services/%s/internal/infrastructure/grpc_client", service),
				fmt.Sprintf("services/%s/internal/transport", service),
				fmt.Sprintf("services/%s/internal/transport/grpc", service),
				fmt.Sprintf("services/%s/internal/transport/http", service),
				fmt.Sprintf("services/%s/configs", service),
			}
		}

		mkdirAll := func(directories []string, wg *sync.WaitGroup) {
			defer wg.Done()

			for _, dir := range directories {
				err := os.MkdirAll(dir, 0755)
				if err != nil {
					fmt.Println("Create error:", err)
				}
				fmt.Printf("Directory created: %s\n", dir)
			}
		}

		_, err := os.Stat("go.mod")
		if err == nil {
			fmt.Println("go.mod file found. Initializing microservices structure...")
		} else if errors.Is(err, os.ErrNotExist) {
			fmt.Println("No go.mod file found.")
			return
		} else {
			fmt.Println("System error:", err)
			return
		}

		var wg sync.WaitGroup

		for _, name := range serviceName {
			if strings.TrimSpace(name) != "" {
				fmt.Printf("Initializing service: %s\n", name)
				wg.Add(1)
				go mkdirAll(generatePaths(name), &wg)
			}
		}

		directories := []string{
			"proto",
			"pkg",
			"pkg/logger",
			"pkg/response",
			"pkg/interceptor",
			"deployments",
			"api-gateway",
			"api-gateway/cmd",
			"api-gateway/internal",
			"api-gateway/internal/handler",
			"api-gateway/internal/middleware",
			"api-gateway/internal/client",
			"api-gateway/internal/aggregator",
			"api-gateway/configs",
		}
		wg.Add(1)
		go mkdirAll(directories, &wg)

		wg.Wait()
		fmt.Println("Microservices structure initialized successfully.")
	},
}

func init() {
	rootCmd.AddCommand(microservicesCmd)

	microservicesCmd.Flags().StringArrayVarP(&serviceName, "service", "s", nil, "Name of the service to initialize")
}
