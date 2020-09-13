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

#### Public vs Private clients.
```golang

// Public clients are only able to access public endpoints.
public := valr.NewPublicClient()

// A normal (or private) client is able to access all endpoints.
private := valr.NewClient("k3y", "s3cr3t")

// You can convert a public client to a private client if ou want to.
private = valr.ToPrivateClient(public, "k3y", "s3cr3t")

// ...or vice versa.
public = valr.ToPublicClient(private)

```

#### Configuring the client.
```golang
client := valr.DefaultClient

client.SetProxyURL("https://proxy.example.com").SetHeader("X-Powered-By", "valr")
```

#### Fetching the Order Book.
```golang
book, err := client.OrderBook(context.Background(), "BTCZAR")
if err != nil {
	log.Fatal(err)
}
```

#### Advanced configuration.
By default, the underlying HTTP client used is a [Snorlax client](https://github.com/nickcorin/snorlax).
Check the documentation to see all the configuration options.
```golang
// You can create your own custom HTTP client.
customClient := snorlax.DefaultClient.SetProxy("https://www.example.com")

// ...and use it for the VALR client.
valrClient := valr.NewClient().SetHTTPClient(customClient)

```

## Contributing
Please feel free to submit issues, fork the repositoy and send pull requests!

## License
This project is licensed under the terms of the MIT license.
