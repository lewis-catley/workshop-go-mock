# Better example

In this example we have altered the code so that we can now inject mocks.

In http-client.go we now create our own `IHTTPClient` interface which contains the same `Get` function signature as `http.Get` the `gomock` library can make use of this later on. We also created a struct has a property of the same type of this `IHTTPClient` interface. So when we call `NewWebClient()` we can return `http.Client{}` as this contains the `Get` function we defined in the `HTTPClient` interface.

An interface has also been created to encapsulate the `GetContent` function so we can also mock this interface for when this function is called in `main.go -> printUrlContent()`

At the bottom of `http-client.go` you can see we use the `go:generate` tool to create mocks by calling `mockgen`.
