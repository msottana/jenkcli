# jenkcli
A simple Jenkins client written in Go

# Docker image
To build the project image use the command:

```docker build . -t jenkcli```

Once built, the program can be started with:

```docker run -v <path/to/config/file>:jenkcli-auth.yaml jenkcli <params>```