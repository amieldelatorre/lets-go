package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// - Go's file server sanitizes request paths by running them through the path.Clean() function which removes any "." and ".." elements to prevent directory traversal
	// - Range requests are also fully supported
	// - The Last-Modified and If-Modified-Since headers are transparently supported. If a file hasn’t changed since the user last requested it,
	// 			then http.FileServer will send a 304 Not Modified status code instead of the file itself.
	// - The Content-Type is automatically set from the file extension using the mime.TypeByExtension() function.
	//			You can add your own custom extensions and content types using the mime.AddExtensionType()
	// - Serves it out of hard disk. But it's likely that http.FileServer will be serving them from RAM because Windows and Unix-based OSs cache recently-used files in RAM

	// Disable directory listing by:
	// Adding a blank index.html to specific directories
	// Or create a custom implementation of http.Filesystem and have it return an os.ErrNotExist for any directories
	// https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
