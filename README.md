# mfr

(Immutable) Map/Filter/Reduce functions (+ friends).

Disclaimer: if you need as much performance as possible, this library is not for you.

# Install
```bash
$ go get github.com/bnert/mfr
```
Note: must have go1.18 installed.

# Example Usage (Map/Filter/Reduce)
```go
package main

import (
  "fmt"
  "github.com/bnert/mfr"
)

type Data struct {
  Value int `json:"value"`
}

func main() {
  d1 := []Data{
    Data{Value: 0},
    Data{Value: 1},
    Data{Value: 2},
    Data{Value: 3},
    Data{Value: 4},
  }

  // Map to a different type
  d2 := mfr.Map[Data, int](d1, func(ctx mfr.Ctx[Data]) int) []int {
    return ctx.Item.Value * 3
  })

  d3 := mfr.Filter[Data](d2, func(ctx mfr.Ctx[Data]) bool) []Data {
    return ctx.Item % 2 == 0
  })

  sum := mfr.Reduce[int, int](d3, 0, func(ctx mfr.Ctx[int], acc int) int {
    return ctx.Current + ctx.Item
  })

  // Prints:
  // Reduced even sum: 18
  fmt.Println("Reduced evens sum:", sum)
}
```
Please see `examples/` and `tests/` directory for other examples of how this
library can be used to write (somewhat) functional go.
