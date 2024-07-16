# Setting Up a New Go Module with gqlgen

## Step 1: Create a New Go Module

1. Create and navigate to a new directory for your project:
    ```bash
    mkdir gqlgen-client2
    cd gqlgen-client2
    ```

2. Initialize a new Go module:
    ```bash
    go mod init github.com/darashevcstbg/gqlgen-client2
    ```

## Step 2: Add gqlgen to Your Project’s tools.go

1. Add gqlgen and introspection to `tools.go`:
    ```bash
    printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
    ```

2. Tidy up the module dependencies:
    ```bash
    go mod tidy
    ```

## Step 3: Initialize gqlgen Config and Generate Models

1. Initialize gqlgen configuration and generate initial models:
    ```bash
    go run github.com/99designs/gqlgen init
    ```

2. Tidy up the module dependencies again:
    ```bash
    go mod tidy
    ```

## Step 4: Modify the GraphQL Schema

1. Open the `graph/schema.graphqls` file.

2. **Add your schema definitions and extend the `Query` and `Mutation` types:**

    ```graphql
    extend type Query {
        # Add your query fields here
    }

    extend type Mutation {
        # Add your mutation fields here
    }
    ```

## Step 5: Create a Custom Code Generator

1. Create a `generate/generate.go` file with the following content:

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

2. Replace the `MODELS_PATH` variable with the path to your client models, e.g., `github.com/darashevcstbg/gqlgen-client2/graph/model`.

## Step 6: Install the Custom gqlgen Library

1. Get the gqlgen library:
    ```bash
    go get github.com/darashevcstbg/gqlgen
    ```

## Step 7: Define the Meetup Schema

1. Create a schema file `graph-lib/meetup.graphqls` with the following content:

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

## Step 8: Update the Resolver

1. Update the `graph/resolver.go` to use the generated file:

    ```go
    package graph

    //go:generate go run ../generate/generate.go

    // This file will not be regenerated automatically.
    // It serves as dependency injection for your app, add any dependencies you require here.

    type Resolver struct{}
    ```

## Step 9: Configure gqlgen

1. Add the new schema file to `gqlgen.yml`:

    ```yaml
    schema:
      - graph/*.graphqls
      - graph-lib/*.graphqls
    ```

## Step 10: Clean Up Old Schema Resolvers

1. Delete the old schema resolver file:
    ```bash
    rm -f graph/schema.resolvers.go
    ```

## Step 11: Generate Code

1. Execute the code generation:
    ```bash
    go generate ./...
    ```

## Step 12: Start the GraphQL Server

1. Run the GraphQL server:
    ```bash
    go run server.go
    ```