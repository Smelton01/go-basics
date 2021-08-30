# HTML Link Parser

## Details

A package that makes it easy to parse an HTML file and extract all of the links (`<a href="">...</a>` tags). For each link we return a data structure that includes both the `href`, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

Links will be nested in different HTML elements, an example is HTML similar to code below.

```html
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
```

Expected Output:

```go
Link{
  Href: "/dog",
  Text: "Something in a span Text not in a span Bold text!",
}
```

## Tests

Set up tests in `link_test.go` and run the following code:

```
go test -v
```
