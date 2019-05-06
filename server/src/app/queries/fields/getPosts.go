package queries

import (
    "context"
    "github.com/graphql-go/graphql"
    "github.com/mongodb/mongo-go-driver/bson"

    "app/data"
    types "app/types"
)

type postStruct struct {
    TITLE string `json:"title"`
    BODY  string `json:"body"`
}

var GetPosts = &graphql.Field {
    Type:        graphql.NewList(types.Post),
    Description: "Get all posts",
    Resolve:     func(params graphql.ResolveParams) (interface{}, error) {

        postCollection := mongo.Client.Database("medium-app").Collection("Posts")
        posts, err := postCollection.Find(context.Background(), nil)
        if err != nil { panic(err) }
        var postsList []postStruct
        for posts.Next(context.Background()) {
            doc := bson.NewDocument()
            err := posts.Decode(doc)
            if err != nil { panic(err) }
            keys, err := doc.Keys(false)
            if err != nil { panic(err) }
            // convert BSON to struct
            post := postStruct{}
            for _, key := range keys {
                keyString := key.String()
                elm, err := doc.Lookup(keyString)
                if err != nil { panic(err) }
                switch (keyString) {
                    case "title":
                    post.TITLE = elm.Value().StringValue()
                    case "body":
                    post.BODY = elm.Value().StringValue()
                    default:
                }
            }
            postsList = append(postsList, post)
        }

        return postsList, nil
    },
}
