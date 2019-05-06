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
        name, _ := params.Args["name"].(string)
        description, _ := params.Args["description"].(string)
        notTodoCollection := mongo.Client.Database("medium-app").Collection("Not_Todos")
        _, err := notTodoCollection.InsertOne(context.Background(), map[string]string{"name": name, "description": description })
        if err != nil { panic(err) }
        return todoStruct{name, description}, nil
    },
}
