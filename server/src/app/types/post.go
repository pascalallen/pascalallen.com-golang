package types

import (
    "github.com/graphql-go/graphql"
)

var Post = graphql.NewObject(graphql.ObjectConfig {
    Name: "Post",
    Fields: graphql.Fields{
        "title": &graphql.Field{
            Type: graphql.String,
        },
        "body": &graphql.Field{
            Type: graphql.String,
        },
    },
})
