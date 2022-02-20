# blog.peernet.org

The Peernet blog is published at https://blog.peernet.org.

## Compilation Guide

Download and install Go from https://golang.org/dl/. Now you are ready to compile it! From the project directory run:

```
go build
```

To compile the binary for Linux on Windows run below commands. For general information on cross-compilation see http://golangcookbook.com/chapters/running/cross-compiling/.

```
set GOOS=linux
go build
```

## Deployment Guide

It requires the compiled server binary, the YAML configuration file and the html folder.

## Reporting Bugs and Feature requests

Please use the GitHub issue tracker for filing any bugs or feature requests.

## Copyright

&copy; 2021 Peernet s.r.o.
