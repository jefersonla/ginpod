# Dependencies
deps: go.mod go.sum
	go mod vendor

# Run the application
run: deps
	air