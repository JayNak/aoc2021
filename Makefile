build: ## Build your project and put the output binary in build
	mkdir -p bin
	GO111MODULE=on go build -mod vendor -o ./bin/aoc2021 ./cmd/aoc2021

run: build
	./bin/aoc2021