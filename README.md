# 🚀 Go-ThreatMatrix

[![Go Reference](https://pkg.go.dev/badge/github.com/khulnasoft/go-threatmatrix.svg)](https://pkg.go.dev/github.com/khulnasoft/go-threatmatrix)
[![Issues](https://img.shields.io/github/issues/khulnasoft/go-threatmatrix?color=blue&logo=github&logoColor=white&style=flat-square)](https://github.com/khulnasoft/go-threatmatrix/issues)
[![License](https://img.shields.io/github/license/khulnasoft/go-threatmatrix?color=blueviolet&logo=open-source-initiative&style=flat-square)](https://github.com/khulnasoft/go-threatmatrix/blob/main/LICENSE)

> A lightweight Go SDK to seamlessly integrate with [ThreatMatrix](https://github.com/khulnasoft/ThreatMatrix) for automation and threat intelligence.

---

## 📘 Table of Contents

* [Overview](#-overview)
* [🚀 Getting Started](#-getting-started)

  * [📦 Prerequisites](#-prerequisites)
  * [⚙️ Installation](#️-installation)
  * [📌 Usage](#-usage)
  * [🧪 Examples](#-examples)
* [🤝 Contribute](#-contribute)
* [📄 License](#-license)
* [🔗 Links](#-links)
* [❓ FAQ](#-faq)

---

## 📖 Overview

**go-threatmatrix** is a powerful and easy-to-use SDK that helps developers interact with the ThreatMatrix API effortlessly. It provides features to manage jobs, users, tags, and more — built with extensibility and developer happiness in mind.

---

## 🚀 Getting Started

### 📦 Prerequisites

* Go 1.17 or higher

### ⚙️ Installation

Install the SDK using `go get`:

```bash
go get github.com/khulnasoft/go-threatmatrix
```

---

### 📌 Usage

To begin using the SDK, import the package and instantiate the client:

```go
import "github.com/khulnasoft/go-threatmatrix/gothreatmatrix"
```

#### 🔐 Create a client:

```go
clientOptions := gothreatmatrix.ThreatMatrixClientOptions{
	Url:         "https://your-threatmatrix-url",
	Token:       "your-api-token",
	Certificate: "optional-cert", // Optional
}

client := gothreatmatrix.NewThreatMatrixClient(&clientOptions, nil)

ctx := context.Background()
jobs, err := client.JobService.List(ctx)
```

#### 🏷️ Create a Tag:

```go
tagOptions := gothreatmatrix.TagParams{
	Label: "NEW TAG",
	Color: "#ffb703",
}

createdTag, err := client.TagService.Create(ctx, tagOptions)
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(createdTag)
}
```

📚 For advanced configuration, refer to the [examples directory](./examples/) and [package docs](https://pkg.go.dev/github.com/khulnasoft/go-threatmatrix).

---

## 🧪 Examples

```go
package main

import (
	"context"
	"fmt"
	"github.com/khulnasoft/go-threatmatrix/gothreatmatrix"
)

func main() {
	clientOptions := gothreatmatrix.ThreatMatrixClientOptions{
		Url:         "https://your-url",
		Token:       "your-api-token",
		Certificate: "your-cert",
	}

	client := gothreatmatrix.NewThreatMatrixClient(&clientOptions, nil)
	ctx := context.Background()

	user, err := client.UserService.Access(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User Details:", *user)
	}
}
```

---

## 🤝 Contribute

We ❤️ contributions! Whether it's a feature request, bug fix, or suggestion — you're welcome to [join our Slack](https://honeynetpublic.slack.com/archives/C01KVGMAKL6) and get involved.

---

## 📄 License

Licensed under the **GNU AFFERO GENERAL PUBLIC LICENSE**. See the [LICENSE](https://github.com/khulnasoft/go-threatmatrix/blob/main/LICENSE) file for details.

---

## 🔗 Links

* 🌐 [ThreatMatrix](https://github.com/khulnasoft/ThreatMatrix)
* 📚 [Documentation](https://threatmatrix.readthedocs.io/en/latest/)
* 🔍 [API Docs](https://threatmatrix.readthedocs.io/en/latest/Redoc.html)
* 💡 [Examples](./examples/)

---

## ❓ FAQ

### 🔑 Generate API Key

You need a valid API key to authenticate with the ThreatMatrix server.

#### v4.0 and Above

1. Login to ThreatMatrix
2. Click your profile in the top right
3. Select `API Access/Sessions`
4. Generate or view your API key

#### Below v4.0

Generate API keys via the admin panel:

* Navigate to **Durin → Auth Tokens** in the ThreatMatrix admin UI
* Create a new key

---

Let me know if you'd like a version with badges aligned horizontally, embedded screenshots, dark theme styling, or monospace improvements.
