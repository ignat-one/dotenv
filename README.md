# Dotenv Loader

A simple Go package to locate and load `.env` files from any directory efficiently.

## Installation

```bash
go get github.com/ignat-one/dotenv
```

## Usage
```go
package main

import (
    "log"
    "github.com/ignat-one/dotenv"
)

func main() {
    err := dotenv.Load()
    if err != nil {
        log.Fatal(err)
    }
    // environment variables are now loaded
}
```

## Usage with options
```go
package main

import (
    "log"
    "github.com/ignat-one/dotenv"
)

func main() {
    err := dotenv.Load(
        dotenv.WithAnchorFile("your-anchor.file"),
        dotenv.WithFilename("your.env"),
    )
    if err != nil {
        log.Fatal(err)
    }
    // environment variables are now loaded
}
```

## Options

* ```WithAnchorFile(filename string)```: Specify a custom anchor file to help locate the ```.env``` file. The default value is ```"go.mod"```

* ```WithFilename(filename string)```: Specify a custom environment filename. The default value is ```".env"```

## Contributing
Feel free to submit issues and enhancement requests.

## License
This project is licensed under the MIT License - see the LICENSE file for details.






