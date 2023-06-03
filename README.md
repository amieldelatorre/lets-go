# lets-go

Based on the book [`Let's Go` by Alex Edwards](https://lets-go.alexedwards.net/).




> There is one more thing that’s really important to point out: all incoming HTTP requests are served in their own goroutine. For busy servers, this means it’s very likely that the code in or called by your handlers will be running concurrently. While this helps make Go blazingly fast, the downside is that you need to be aware of (and protect against) race conditions when accessing shared resources from your handlers.[^1]

[^1]: Chapter 2.9 Requests are handled concurrently