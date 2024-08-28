# TODO cli app

<video width="320" height="240" controls>
  <source src="/docs/preview.webm" type="video/webm">
  Your browser does not support the video tag.
</video>


### Requirements

- Golang 1.22.6

### How to install locally

```bash
cd ./todo-cli 
go build -o tasks 
sudo ln -s $(pwd)/tasks /usr/local/bin/tasks # Create a symlink to the binary
```

### Using make to compile and install

```bash
make run
```

### How to use cli app

Info about the task cli app
```bash
tasks --help
```