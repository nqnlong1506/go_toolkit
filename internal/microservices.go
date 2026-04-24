package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func MicroservicesInit() error {
	usernameCmd := exec.Command("git", "config", "user.name")
	usernameOutput, err := usernameCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error getting git username:", err)
		fmt.Println("Enter username: ")
	} else {
		fmt.Printf("Enter username (default: %s): ", strings.ReplaceAll(string(usernameOutput), "\n", ""))
	}
	var username string
	fmt.Scanln(&username)
	if strings.TrimSpace(username) == "" {
		username = strings.TrimSpace(string(usernameOutput))
	}

	var projectName string
	fmt.Printf("Enter project name: ")
	fmt.Scanln(&projectName)
	for {
		if strings.TrimSpace(projectName) != "" {
			break
		}

		fmt.Print("Project name cannot be empty. Please enter a valid project name: ")
		fmt.Scanln(&projectName)
	}

	moduleName := fmt.Sprintf("github.com/%s/%s", username, projectName)
	fmt.Printf("Enter module name (default: %s): ", moduleName)
	fmt.Scanln(&moduleName)

	cmdInit := exec.Command("go", "mod", "init", moduleName)
	cmdInit.Stdout = os.Stdout
	cmdInit.Stderr = os.Stderr
	err = cmdInit.Run()
	if err != nil {
		fmt.Println("Error initializing go module:", err)

		return err
	}

	fmt.Printf("Go module initialized with name: %s\n", moduleName)
	return nil
}

func MicroservicesGenerateFolderStructure(serviceNames []string) error {
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
		return errors.New("go.mod file not found. Please run 'go mod init <MODNAME>' first")
	} else {
		fmt.Println("System error:", err)
		return err
	}

	var wg sync.WaitGroup

	for _, name := range serviceNames {
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
	return nil
}

func MicroservicesGenerateDockerfiles(serviceNames []string) {
	targets := []string{
		"api-gateway",
	}
	for _, name := range serviceNames {
		targets = append(targets, fmt.Sprintf("services/%s", name))
	}

	for _, dir := range targets {
		file, err := os.Stat(dir)
		if err != nil || !file.IsDir() {
			fmt.Printf("Directory not found or not a directory: %s\n", dir)
			continue
		} else {
			fmt.Println(dir, "is dir")
		}
	}

	fmt.Println("Generating Dockerfiles...")
}
