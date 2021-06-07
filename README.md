# httpshopify

An HTTP implementation of the [shopify](https://github.com/MOHC-LTD/shopify) package.

## Contents

- [Github authentication](#github-authentication)
- [Installation](#installation)
- [Usage](#usage)
- [Docker](#docker)
- [How to contribute](#how-to-contribute)

## Github authentication

As this is a private Github repository, you will need to set up private go modules.

First set the `GOPRIVATE` go environment variable.

```sh
go env -w GOPRIVATE=github.com/MOHC-LTD
```

Generate a [Github personal access token](https://github.com/settings/tokens), and set up
global Github authentication on your machine

```sh
git config --global url."https://${username}:${access_token}@github.com".insteadOf "https://github.com"
```

## Installation

Install the module using

```sh
go get -u github.com/MOHC-LTD/httpshopify
```

## Usage

If directly talking to a Shopify API then use the `NewShop` constructor. This
contructor will automatically generate the URL of the shop from the shop name.

```go
func main() {
    shop := httpshopify.NewShop("shop-name", "shopify-access-token")

    order, err := shop.Orders().Get("184190283")
}
```

If talking to a custom shopify API, specify the entire URL with the `NewCustomShop` contructor.

```go
func main() {
    shop := httpshopify.NewCustomShop("https://myhost.com", "shopify-access-token")

    order, err := shop.Orders().Get("184190283")
}
```

## Docker

To build applications that consuming this module using docker, you will need to allow the docker container to authenticate with Github.

Do this by adding the following lines to your Dockerfile.

```sh
ARG authToken

RUN go env -w GOPRIVATE=github.com/MOHC-LTD

RUN apk add git

RUN git config --global url."https://golang:$authToken@github.com".insteadOf "https://github.com"
```

Then, when building your container, set the docker argument `authToken` to the value of your Github access token.

## How to contribute

Something missing or not working as expected? See our [contribution guide](./CONTRIBUTING.md).
