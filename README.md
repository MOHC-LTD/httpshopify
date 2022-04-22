# httpshopify

An HTTP implementation of the [shopify](https://github.com/MOHC-LTD/shopify) package.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [How to contribute](#how-to-contribute)

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

You can also specify a Shopify plus shop which will used the increased rate limit.

```go
func main() {
    shop := httpshopify.NewPlusShop("shop-name", "shopify-access-token")

    order, err := shop.Orders().Get("184190283")
}
```

If talking to a custom shopify API, specify the entire URL with the `NewCustomShop` contructor.

```go
func main() {
    shop := httpshopify.NewCustomShop("https://myhost.com", "shopify-access-token", shopify.IsPlus)

    order, err := shop.Orders().Get("184190283")
}
```

## How to contribute

Something missing or not working as expected? See our [contribution guide](./CONTRIBUTING.md).

test change
