build:
	go build -o lio-toolkit && sudo cp lio-toolkit /usr/local/bin/
complement:
	lio-toolkit completion bash | sudo tee /usr/share/bash-completion/completions/lio-toolkit > /dev/null