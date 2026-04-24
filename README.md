# Ben's go lib

This is an assortment of useful stuff for go.

## errors

The `errors` package contains an implementation of the `error` interface which you can use
in constants, and also implements pass-through functions for the rest of the standard
`errors` package, making it a drop-in replacement for the standard errors package.

Thanks to Jonathan Hall of [boldlygo.tech](https://boldlygo.tech/archive/2025-02-12-constant-errors/)
for the basic idea. :)

- Catalan blog post: <https://benjaminnolan.cat/go-error-consts.html>
- English blog post: <https://benjaminnolan.dev/go-error-consts.html>

### `errors.Error` usage examples

```go
package myamazingcode

import (
    "github.com/benjaminnolan/lib/errors"
)

const (
    ErrMyAmazingCodeError1 errors.Error = "this is the first error"
    ErrMyAmazingCodeError2 errors.Error = "this is the second error"
)
```
