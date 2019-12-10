### Description

This Go library is a client for the MeiliSearch Database. 

Currently the whole API is not yet implemented. This is not yet usable in production.

### Dependencies

- github.com/pkg/errors for error wrapping

### How to use it

```go
package main

import (
    "github.com/alexisvisco/meilisearch-go"
    "fmt"
)

func main() {

    // meilisearch package will maybe be renamed into ms ?

    client := meilisearch.NewClient(meilisearch.Config{
        Host: "http://localhost:7700",
    })
    
    index, err := client.Indexes().Create(meilisearch.CreateIndexRequest{
        Name: "Meilimelo",
        Schema: meilisearch.Schema{
            "id": {
                meilisearch.SchemaAttributesIdentifier,
                meilisearch.SchemaAttributesIndexed, 
                meilisearch.SchemaAttributesDisplayed, 
            },
        },
    })
    
    // advanced error handling  err.(*meilisearch.Error)
    // do stuff with index and err
    
    if err != nil {
        panic(err)
    }
    
    fmt.Println(index)
}
```

### Stable part of this library

- Indexes https://docs.meilisearch.com/references/indexes.html 
- Documents https://docs.meilisearch.com/references/documents.html 