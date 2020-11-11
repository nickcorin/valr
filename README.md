<p align="center">
<h1 align="center">VALR</h1>
<p align="center">Unofficial Go Client for VALR</p>
<p align="center">This is currently a work-in-progress and not fully tested.</p>
</p>
<p align="center">
<p align="center"><a href="https://github.com/nickcorin/valr/actions?query=workflow%3AGo"><img src="https://github.com/nickcorin/valr/workflows/Go/badge.svg?branch=master" alt="Build Status"></a> <a href="https://goreportcard.com/report/github.com/nickcorin/valr"><img src="https://goreportcard.com/badge/github.com/nickcorin/valr?style=flat-square" alt="Go Report Card"></a> <a href="http://godoc.org/github.com/nickcorin/valr"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square" alt="GoDoc"></a> <a href="LICENSE"><img src="https://img.shields.io/github/license/nickcorin/valr" alt="License"></a></p>
</p>
<p align="center">
<img src="/images/valr.png" />
</p>

## Installation

To install `valr`, use `go get`:
```
go get github.com/nickcorin/valr
```

Import the `valr` package into your code:
```golang
package main

import "github.com/nickcorin/valr"

func main() {
	client := valr.DefaultClient
}
```

## Usage

#### Using the DefaultClient.
```golang
// The default (public) client can access all public endpoints without providing authentication.
ctx := context.Background()
book, err := valr.DefaultClient.OrderBook(ctx, "BTCZAR")
if err != nil {
  log.Fatal(err)
}
```

#### Accessing authenticated endpoints.
```golang
// To access authentiated endpoints, you need to construct a (private) client.
client := valr.NewClient("my-api-key", "my-api-secret")

ctx := context.Background()
orderID, err := client.LimitOrder(ctx, valr.LimitOrderRequest{
  CustomerOrderID:  "1234",
  Pair:             "BTCZAR",
  PostOnly:         true,
  Price:            "200000",
  Quantity:         "0.100000",
  Side:             "SELL",
})
if err != nil {
  log.Fatal(err)
}
```

#### Public vs Private clients.
```golang

// Public clients are only able to access public endpoints.
public := valr.NewPublicClient()

// A normal (or private) client is able to access all endpoints.
private := valr.NewClient("my-api-key", "my-api-secret")

// You can convert a public client to a private client if ou want to.
private = valr.ToPrivateClient(public, "my-api-key", "my-api-secret")

// ...or vice versa.
public = valr.ToPublicClient(private)

```

#### Fetching the Order Book.
```golang
ctx := context.Background()
book, err := client.OrderBook(ctx, "BTCZAR")
if err != nil {
	log.Fatal(err)
}
```

## Contributing
Please feel free to submit issues, fork the repositoy and send pull requests!

## License
This project is licensed under the terms of the MIT license.
