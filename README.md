# lets-go

Based on the book [`Let's Go` by Alex Edwards](https://lets-go.alexedwards.net/).

- Go's file server sanitizes request paths by running them through the path.Clean() function which removes any "." and ".." elements to prevent directory traversal
- Range requests are also fully supported
- The Last-Modified and If-Modified-Since headers are transparently supported. If a file hasn’t changed since the user last requested it, then http.FileServer will send a 304 Not Modified status code instead of the file itself.
- The Content-Type is automatically set from the file extension using the mime.TypeByExtension() function. You can add your own custom extensions and content types using the mime.AddExtensionType()
- Serves it out of hard disk. But it's likely that http.FileServer will be serving them from RAM because Windows and Unix-based OSs cache recently-used files in RAM

- Disable directory listing by:
- Adding a blank index.html to specific directories
- Or create a custom implementation of http.Filesystem and have it return an os.ErrNotExist for any directories
- https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings

> There is one more thing that’s really important to point out: all incoming HTTP requests are served in their own goroutine. For busy servers, this means it’s very likely that the code in or called by your handlers will be running concurrently. While this helps make Go blazingly fast, the downside is that you need to be aware of (and protect against) race conditions when accessing shared resources from your handlers.[^1]

> Closures for dependency injection by moving it out into a separate file.[^2]

> We could redirect the stdout and stderr streams to on-disk files when starting the application
> $ go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log

[^1]: Chapter 2.9 Requests are handled concurrently
[^2]: Chapter 3.3 Closures for dependency injection https://gist.github.com/alexedwards/5cd712192b4831058b21