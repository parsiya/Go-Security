# spf13/cobra examples

## cobra generator installation
The `cobra` generator is not installed with `go get` (on Windows at least), to build and install:

1. `go get github.com\spf13\viper`.
2. `go get github.com/mitchellh/go-homedir`.
3. Navigate to `GOPATH\src\github.com\spf13\cobra\cobra` (note the second `cobra` directory).
4. `go install`.

## Generating sample app
Run `cobra init github.com/blah/whatever` to create the structure for an app. The directory will be created under `src`. So in this case I ran `cobra init Go-Security/cli-package-tests/spf13-cobra/app`.

If not specified, the default LICENSE is Apache v2. Replace it with your own license if needed.

We can configure author and license with flags: `cobra init app -l MIT -a "Parsia"`.

### main.go
Main just calls `cmd.Execute()` which is in `app/cmd/root.go`.

### cmd/root.go
Inside `root.go` we can define what the application does when nothing is passed to it.

``` go
var rootCmd = &cobra.Command{
    Use:   "cobra-example",
    Short: "Sample app to learn Cobra",
    // Long is printed with -h
    Long: `Sample app to learn Cobra.
Second line.`,
    // Uncomment the following line if your bare application
    // has an action associated with it:
    // Here we can pass our own function for noArgs.
    Run: noArgs,
}

// noArgs is run if program is run without arguments.
func noArgs(cmd *cobra.Command, args []string) {
    fmt.Println("please provide some arguments")

    // Print usage string
    fmt.Println(cmd.UsageString())
}
```

Cobra by default adds support to read info from cfg files using Viper.

### cmd/cmd1.go
This will be created if we do `cobra add cmd1`. We can set the description and other info for the cmd here. We can create aliases here by creating a string slice.

``` go
// cmd1Cmd represents the cmd1 command
var cmd1Cmd = &cobra.Command{
    Use:   "cmd1",
    Short: "Short description of cmd1",
    Long:  `Long description of cmd1`,
    // cmd1 aliases
    Aliases: []string{"cmd11", "cmd12"},
    Run:     cmd1Executor,
}

// cmd1Executor is run when cmd1 is called.
func cmd1Executor(cmd *cobra.Command, args []string) {
    // Here we can see what command was run and also any arguments.
    fmt.Println("cmd1 called")
    // Print arguments after cmd1
    fmt.Println("arguments", args)
}
```

We can call `cmd` or any of the aliases and access the arguments after it.

```
$ go run main.go cmd11
cmd1 called

$ go run main.go cmd12
cmd1 called

$ go run main.go cmd1 arg1 arg2
cmd1 called
arguments [arg1 arg2]

```

### cmd2.go
We create `cmd2` and assign `cmd1` as its parent. `cobra add cmd2 -p "cmd1Cmd"`. Note `cmd1Cmd` is correct and not `cmd1`. To get the name of the command, go to `cmd1.go` and look inside `init()`:

``` go
func init() {
    rootCmd.AddCommand(cmd1Cmd)
}
```

If this happens, modify the `init` function inside `cmd2.go`:

``` go
func init() {
    cmd1Cmd.AddCommand(cmd2Cmd)
}
```

Now we can run `cmd2` as a subcommand of `cmd1`.

```
$ go run main.go cmd1 cmd2
cmd2 called
```

And it shows up in `cmd1`'s help.

```
$ go run main.go cmd1 -h
Long description of cmd1

Usage:
  cobra-example cmd1 [flags]
  cobra-example cmd1 [command]

Aliases:
  cmd1, cmd11, cmd12

Available Commands:
  cmd2        Short description of cmd2

Flags:
  -h, --help   help for cmd1

Global Flags:
      --config string   config file (default is $HOME/.app.yaml)

Use "cobra-example cmd1 [command] --help" for more information about a command.
```

