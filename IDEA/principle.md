# Draft

## 1. To see whether the site accepts partial-requests

Usually a site tells you whether it accepts partial-requests, with the "Accept-Ranges" head.

In case the header `Accept-Ranges: <unit>` presents, the browser may try to resume an intererupted download, thus making multi-threaded downloading possible.

For example
> Accept-Ranges: bytes // ok
  Accept-Ranges: none // not ok ;)

So, we need to send a Header request to the site:

```go
res, err := http.Head(target_url)
if err != nil {
    return err
}

if res.StatusCode == http.StatusOK && res.Header.get("Accept-Ranges") == "bytes" {
    // blahblah
}
```

---

## 2. How to make partial requests

---

The "Range" HTTP request header indicates the part of a document that the server should return. You may tell the server multiple ranges you require in the `Rnage` Header.

Several parts can be requested with one `Range` header at once, and the server may send back these ranges in a multipart document. If the server sends back ranges, it uses the `206 Partial Content` for the response. If the ranges are invalid, the server returns the `416 Range Not Satisfiable` error. The server can also ignore the Range header and return the whole document with a `200` status code.

The `Range` Header should be written like this

```plain-text
Range: <unit>=<range-start>-
Range: <unit>=<range-start>-<range-end>
Range: <unit>=<range-start>-<range-end>, <range-start>-<range-end>
```

`<unit>`

The unit in which ranges are specified. This is usually `bytes`.

`<range-start> / <range-end>`

An integer in the given unit indicating the beginning and the end of the request range. The value of "range-end" is optional and, if omitted, the end of the document is taken as the end of the range.

//`<suffix-length>`
//
//An integer in the given unit indicating the number of units at the end of the file to return.

### Examples

requesting three ranges from the file:

`Range: bytes=200-1000, 2000-6576, 19000-`

refer to the Mozilla References for more [documents of Range](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Range)
