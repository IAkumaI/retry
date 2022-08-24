# retry

Simple library for retry function on error

## Installation

`go get github.com/IAkumaI/retry`

## Usage

```go
package main

import (
	"errors"
	"github.com/IAkumaI/retry"
	"log"
)

func main() {
	// Function will be called until no errors up to 3 times
	// If after 3 runs there is still error,
	// this error will be returned by Do method
	err := retry.Do(3, func(retryCount int) error {
		log.Println("retryCount", retryCount)

		if retryCount < 2 {
			return errors.New("test error")
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
```