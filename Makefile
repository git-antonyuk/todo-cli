.PHONY: run

run:
		go build -o tasks
		    if [ ! -L /usr/local/bin/tasks ]; then sudo ln -s $(PWD)/tasks /usr/local/bin/tasks; fi
		tasks --help