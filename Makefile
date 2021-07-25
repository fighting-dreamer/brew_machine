.PHONY: all
all: clean compile

APP=brew_machine
APP_EXECUTABLE="./out/$(APP)"

clean: ## remove executable
	go mod tidy -v
	rm -f $(APP_EXECUTABLE)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	go get

test: clean go-get ## test the code
	go test ./...

compile: go-get ## build the executable
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)