# Creation steps

## Create a new go module
```bash
mkdir gqlgen-client2
cd gqlgen-client2
go mod init github.com/darashevcstbg/gqlgen-client2
```

## Add github.com/99designs/gqlgen to your project’s tools.go

```bash
printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
go mod tidy
```

## Initialise gqlgen config and generate models
```bash
go run github.com/99designs/gqlgen init
go mod tidy
```

## Modify the graph/schema.graphqls file

Add your schema to the graph/schema.graphqls file, also extend the Query and Mutation types

## Create generate/generate.go file, replace the MODELS_PATH variable with the path to client models (e.g. github.com/darashevcstbg/gqlgen-client2/graph/model)

```go
package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/darashevcstbg/gqlgen/plugins/gqlgen_plugin"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	err = api.Generate(cfg,
		api.AddPlugin(gqlgen_plugin.New("graph/meetup.resolvers.go", "MODELS_PATH")),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
```

## Get the gqlgen library

```bash
go get github.com/darashevcstbg/gqlgen
```

## Create the Meetup schema graph-lib/meetup.graphqls

```graphql
type Meetup {
    id: ID!
    name: String!
    description: String!
    user: User!
}

input NewMeetup {
    name: String!
    description: String!
}

type User {
    id: ID!
    username: String!
    email: String!
    meetups: [Meetup!]!
}

type Query {
    meetups: [Meetup!]!
    users: [User!]!
}

input NewUser {
    username: String!
    email: String!
}
type Mutation {
    createMeetup(input: NewMeetup!): Meetup!
    createUser(input: NewUser!): User!
}
```

### Update the graph/resolver.go to use the generate/generated.go file

```go
package graph

//go:generate go run ../generate/generate.go

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}
```

### Add the graph-lib/*.graphqls schemas to the gqlgen.yml file

```yaml
schema:
  - graph/*.graphqls
  - graph-lib/*.graphqls
```

### Delete the old schema resolver 
```bash
rm -f graph/schema.resolvers.go
```

### Execute go generate

```bash
go generate ./...
```

## Start the graphql server

```bash
go run server.go
```