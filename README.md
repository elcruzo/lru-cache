# Go LRU Cache

![Go Version](https://img.shields.io/badge/Go-v1.17-blue)
![License](https://img.shields.io/badge/License-MIT-green)

A simple thread-safe LRU (Least Recently Used) cache implementation in Go.

## Table of Contents

- [Go LRU Cache](#go-lru-cache)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Usage](#usage)
  - [Examples](#examples)
  - [Contributing](#contributing)
  - [License](#license)

## Introduction

This Go LRU Cache is designed to provide an easy-to-use, thread-safe, and efficient caching mechanism for your Go applications. It uses a double-linked list to maintain the order of recently accessed items, ensuring that the least recently used items are removed when the cache reaches its capacity.

## Features

- Simple and efficient LRU caching.
- Thread-safe operations using mutex locks.
- Customizable cache capacity.
- Get and Put operations for key-value pairs.
- Easy integration into your Go projects.

## Getting Started

### Installation

To use this LRU Cache in your Go project, you can use Go modules. Import it as follows:

```go
import "github.com/elcruzo/mylibrary/lrucache"
```

### Usage

```go

cache := lrucache.NewLRUCache(100)


cache.Put("key1", "value1")
cache.Put("key2", "value2")


value := cache.Get("key1")
fmt.Println(value)
```

## Examples

For more usage examples and advanced features, check the [examples](examples/) directory in this repository.

## Contributing

Contributions are welcome! If you find a bug or have an enhancement in mind, please open an issue or submit a pull request. For more details, see [CONTRIBUTING.md](CONTRIBUTING.md).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
