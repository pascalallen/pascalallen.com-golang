package mutations

import (
    "github.com/graphql-go/graphql"
    "context"
    "app/data"
    types "app/types"
)

type postStruct struct {
    TITLE  string `json:"title"`
    BODY   string `json:"body"`
}

var CreatePost = &graphql.Field {
    Type:        types.Post,
    Description: "Create a post",
    Args: graphql.FieldConfigArgument{
        "title": &graphql.ArgumentConfig{
            Type: graphql.String,
        },
        "body": &graphql.ArgumentConfig{
            Type: graphql.String,
        },
    },

    Resolve: func(params graphql.ResolveParams) (interface{}, error) {
        // get our params
        title, _ := params.Args["title"].(string)
        body, _ := params.Args["body"].(string)
        postCollection := mysql.Client.Database("pascalallen.com").Collection("Posts")
        _, err := postCollection.InsertOne(context.Background(), map[string]string{"title": title, "body": body })
        if err != nil { panic(err) }
        return postStruct{title, body}, nil
    },
}
