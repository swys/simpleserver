# simpleserver
simple http server written in Go with basic routing and logging. Useful as a starting point to create a REST server. Provides basic boiler plate to wire up routing and logging. Other dependancies can be injected and passed around (like a database, for example) by adding to the `Handlers` struct (in routes/routes.go), and then adding the setup code in main.go and passing it to the `routes.NewHandlers` constructor.

This also includes the files needed to setup vscode remote dev container environment. More details here : https://code.visualstudio.com/docs/remote/containers

# License
MIT

