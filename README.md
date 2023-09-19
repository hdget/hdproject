## HD Project template

### Quick Start

##### 1. Install `gonew` tool

```
go install golang.org/x/tools/cmd/gonew@latest
```

##### 2. If new project under China, please configure `GOPROXY`

##### 3. New project using `gonew`

```
gonew github.com/hdget/hdproject example.com/myproject
```

### Directory
- cmd: all cli command utilities
- conf: project configuration
- g: project global variable
- pkg: project business logic