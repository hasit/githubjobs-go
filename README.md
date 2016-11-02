# githubjobs-go

[![Go Report Card](https://goreportcard.com/badge/github.com/hasit/githubjobs-go)](https://goreportcard.com/report/github.com/hasit/githubjobs-go)
[![GoDoc](https://godoc.org/github.com/hasit/githubjobs-go?status.svg)](https://godoc.org/github.com/hasit/githubjobs-go)

Github Jobs SDK for Go.

## Install

```bash
go get github.com/hasit/githubjobs-go
```

## Usage

Assuming that you have installed `githubjobs` package correctly, simply import it to start using. Head over to the [githubjobs GoDoc page](https://godoc.org/github.com/hasit/githubjobs-go) for more reference.

```go
package main

import "github.com/hasit/githubjobs-go"

...
```

You can play with examples in the [examples folder](https://github.com/hasit/githubjobs-go/tree/master/examples) as well. 

### Get a single position by ID

```go
package main

import (
	"fmt"
	"log"

	"github.com/hasit/githubjobs-go"
)

func main() {
	p, err := githubjobs.GetPositionByID("<simple-sample>")
	if err != nil {
		log.Fatal(err)
	}
    // p now contains the job position with <simple-sample> ID.
	fmt.Println(p)
}
```

### Get a list of positions

**By description**

Call `githubjobs.GetPositions()` with empty `location` parameter.

```go
package main

import (
	"fmt"
	"log"

	"github.com/hasit/githubjobs-go"
)

func main() {
	p, err := githubjobs.GetPositions("go", "", false)
	if err != nil {
		log.Fatal(err)
	}
    // p now contains all positions with "go" in their descriptions. 
	fmt.Println(p[0])
}
```

**By location**

Call `githubjobs.GetPositions()` with empty `description` parameter.

```go
package main

import (
	"fmt"
	"log"

	"github.com/hasit/githubjobs-go"
)

func main() {
	p, err := githubjobs.GetPositions("", "seattle", true)
	if err != nil {
		log.Fatal(err)
	}
    // p now contains all positions with "seattle" as their location. 
	fmt.Println(p[0])
}
```

**By geographical coordinates**

Call `githubjobs.GetPositionsByCoordinates` with the coordinates (latitude and longitude) of location. It is important to note that the latitude and longitude values are in decimal degrees. You can use a service like [LatLong.net](http://www.latlong.net) for finding the coordinates of a location of your choice.

```go
package main

import (
	"fmt"
	"log"

	"github.com/hasit/githubjobs-go"
)

func main() {
    // geographical coordinates of Seattle are 47.6062100° and -122.3320700° in decimal degrees.
	p, err := githubjobs.GetPositionsByCoordinates("47.6062100", "-122.3320700")
	if err != nil {
		log.Fatal(err)
	}
    // p now contains 
	fmt.Println(p)
}
```

## Contributing

Questions, comments, bug reports, and pull requests are all welcome.

## Author

[Hasit Mistry](https://github.com/hasit)