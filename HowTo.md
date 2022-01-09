# How to initialize go env
1. First install go:
```bash
brew install go
```

If you want to check the documentation about the *fmt* module, you can do the next command:
```bash
go doc fmt
go doc fmt.Println
```

2. VisualStudio Codium
The best way is to install Visual Studio Codium and then install the Go extension.
After that use CTRL(command) + SHIFT + P to select all of the choices and install them.

3. init
go to the folder where you want to initialize your go app and do next command:
```bash
go mod init github.com/ognjen-it/go-app
```

4. Declaring Variables:
could declare in two or one line:

```go
var i int
i = 42
```
or
```go
var i int = 42
```
or leave go to set init/float/string..
```go
i := 42
```

The const cannot be reassign, for example:
```go
const pi = 3.1415
fmt.Println(pi)
pi = 1.2
```

With const we can set multi const:
```go
const (
    first = 1
    second = "second"
)
```

iota - every time we use it then it increases exponentially. For example:

```go
const (
	first = iota
	second = iota
	third   // here you can see that third inherit last defined const
)

fmt.Println(2 << first)
fmt.Println(2 << second)
fmt.Println(2 << third)

```