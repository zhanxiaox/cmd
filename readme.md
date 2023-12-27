# Overview

Cmd is a simple cmd tool for golang

# Installing

Using cmd is easy.

```go
import "github.com/zhanxiaox/cmd"
```

```bash
go mod tidy
```

# Usage

First import package

```go
import "github.com/zhanxiaox/cmd"
```

Create a app

```go
app := cmd.New()
```

Add a command

```go
app.AddCommand(cmd.Command{Name:"help",Desc:"Print some message",Excute:func(this cmd.Command) error{
	// this is default help for app
	// or you can do some other staff
	return app.DefaultHelp();
}})
```

# License
