# Float Breakdown

This was just a small project to learn about how floats are done under the hood. I only did float32, but it can easily be brought up to float64.

The main part that intrigued me was how the mantissa (fraction) bits were used to calculate the final value. I used Golang's native floats to compute the bits, before I break it down.

All implementation was source from: https://en.wikipedia.org/wiki/IEEE_754

```bash
$ float-breakdown 0.15625
Value to breakdown: 0.156250
    Binary:    0.15625 0b00111110001000000000000000000000
      Sign:           0 b0
  Exponent:          -3 0b1111100
  Mantissa:     2097152 0b1000000000000000000000

Throw this into google to get your number.
(-1)^0 * 1.25 * 2^(-3)
```

# Go Playground

You don't need to clone this library to give it a go. You can use a go [playground online](playground). Just import `"github.com/Emyrk/float-breakdown/breakdown"`

```go
package main

import (
	"fmt"

	"github.com/Emyrk/float-breakdown/breakdown"
)

func main() {
	var value float32 = 0.15625
	fmt.Printf("Value to breakdown: %f\n", value)
	b := breakdown.Float32(value)
	fmt.Println(b)
}
```


[playground]: (https://goplay.tools/snippet/KnfD5kUGTzA)