# URL Shortener

## Details

An [http.Handler](https://golang.org/pkg/net/http/#Handler) looks at the path of any incoming web request and determine if it should redirect the user to a new page.

For instance, if we have a redirect setup for `/dogs` to `https://www.somesite.com/a-story-about-dogs` we would look for any incoming web requests with the path `/dogs` and redirect them.

The `MapHandler` function uses a predefined map to map the path to the corresponding URL.

The `YAMLHandler` and `JSONHandler` functions take a file (`-file <filename>`) and parses the file for the url paths ans corresponding URLs to redirect users as appropriate. The file parsed depends on the extension of the provided file `*.ext`.

The `DBHandler` function checks the database for the provided paths and redirects as appropriate.
Set the flag `-newdb true` to initialize a new database.

## Usage

```
go run main/main.go -newdb true
```
