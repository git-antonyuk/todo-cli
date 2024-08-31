.PHONY: run

run:
		rm -rf ./data/db.csv
		go build -o tasks
		    if [ ! -L /usr/local/bin/tasks ]; then sudo ln -s $(PWD)/tasks /usr/local/bin/tasks; fi
		tasks --help

.PHONY: test

test:
		go test ./cmd/tasks -v

.PHONY: coverage

coverage:
		mkdir -p ./coverage
		go test -cover ./cmd/tasks
		go test -coverprofile=./coverage/coverage.out ./cmd/tasks
		go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
		rm ./coverage/coverage.out