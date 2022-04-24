# telnetClient

a *even easier* telnet client for command servers

[pkg.go.dev](//pkg.go.dev/github.com/ev2-1/telnetClient)

### Connect to server:

```go
c, err := telnetClient.NewController("<destination>")
if err != nil {
	fmt.Println(err)
}
```

### Send a command

```go
resp, err := c.Exec("<command>")
if err != nil {
	fmt.Println(err)
}

fmt.Println(resp)
```

### End connection

```go
err := c.Close()
if err != nil {
	fmt.Println(err)
}
