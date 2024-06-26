# go-intelx
[![GitHub issues](https://img.shields.io/github/issues/khulnasoft/go-intelx?style=plastic)](https://github.com/khulnasoft/go-intelx/issues)
[![GitHub license](https://img.shields.io/github/license/khulnasoft/go-intelx?style=plastic)](https://github.com/khulnasoft/go-intelx/blob/main/LICENSE)

go-intelx is a client library/SDK that allows developers to easily automate and integrate [IntelX](https://github.com/khulnasoft/IntelX) with their own set of tools!

<!-- omit in toc -->
# Table of Contents
- [go-intelx](#go-intelx)
- [Getting Started](#getting-started)
	- [Pre requisites](#pre-requisites)
	- [Installation](#installation)
	- [Usage](#usage)
	- [Examples](#examples)
- [Contribute](#contribute)
- [License](#liscence)
- [Links](#links)
- [FAQ](#faq)
	- [Generate API key](#generate-api-key)
		- [v4.0 and above](#v40-and-above)
		- [v4.0 below](#v40-below)



# Getting Started

## Pre requisites
- Go 1.17+

## Installation
Use go get to retrieve the SDK to add it to your GOPATH workspace, or project's Go module dependencies.

```bash
$ go get github.com/khulnasoft/go-intelx
```

## Usage
This library was built with ease of use in mind! Here are some quick examples to get you started. If you need more example you can go to the [examples directory](./examples/)

To start using the go-intelx library you first need to import it:
```
import "github.com/khulnasoft/go-intelx/gointelx"
```
Construct a new `IntelXClient`, then use the various services to easily access different parts of Intelx's REST API. Here's an example of getting all jobs:

```Go
clientOptions := gointelx.IntelXClientOptions{
	Url:         "your-cool-URL-goes-here",
	Token:       "your-super-secret-token-goes-here",
	// This is optional
	Certificate: "your-optional-certificate-goes-here",
}

intelx := gointelx.NewIntelXClient(
	&clientOptions,
	nil
)

ctx := context.Background()

// returns *[]Jobs or an IntelXError!
jobs, err := intelx.JobService.List(ctx)
```
For easy configuration and set up we opted for `options` structs. Where we can customize the client API or service endpoint to our liking! For more information go [here](). Here's a quick example!

```Go
// ...Making the client and context!

tagOptions = gointelx.TagParams{
  Label: "NEW TAG",
  Color: "#ffb703",
}

createdTag, err := intelx.TagService.Create(ctx, tagOptions)
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(createdTag)
}
```
## Examples
The [examples](./examples/) directory contains a couple for clear examples, of which one is partially listed here as well:

```Go
package main

import (
	"fmt"

	"github.com/khulnasoft/go-intelx/gointelx"
)

func main(){
	intelxOptions := gointelx.IntelXClientOptions{
		Url:         "your-cool-url-goes-here",
		Token:       "your-super-secret-token-goes-here",
		Certificate: "your-optional-certificate-goes-here",
	}	

	client := gointelx.NewIntelXClient(
		&intelxOptions,
		nil,
	)

	ctx := context.Background()

	// Get User details!
	user, err := client.UserService.Access(ctx)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	} else {
		fmt.Println("USER Details")
		fmt.Println(*user)
	}
}

```
For complete usage of go-intelx, see the full [package docs](https://pkg.go.dev/github.com/khulnasoft/go-intelx).

# Contribute
If you want to follow the updates, discuss, contribute, or just chat then please join our [slack](https://honeynetpublic.slack.com/archives/C01KVGMAKL6) channel we'd love to hear your feedback!

# License
Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE.

# Links
- [Intelx](https://github.com/khulnasoft/IntelX)
- [Documentation](https://intelx.readthedocs.io/en/latest/)
- [API documentation](https://intelx.readthedocs.io/en/latest/Redoc.html)
- [Examples](./examples/)

# FAQ
## Generate API key
You need a valid API key to interact with the IntelX server.
### v4.0 and above
You can get an API by doing the following:
1. Log / Signin into intelx
2. At the upper right click on your profile from the drop down select `API Access/ Sessions`
3. Then generate an API key or see it!

### v4.0 below
Keys should be created from the admin interface of [IntelX](https://github.com/khulnasoft/intelx): you have to go in the *Durin* section (click on `Auth tokens`) and generate a key there.
