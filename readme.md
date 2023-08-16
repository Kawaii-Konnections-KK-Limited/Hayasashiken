# Hayasashiken
simple speed test

# foreignusage

This package provides functions for testing network latency of links stored in a database. 

## Usage

```go
import "github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/foreignusage"
```

The main function is `GetTestResults`, which takes a slice of `models.Links` from a database and returns latency test results.

`models.Links` is defined as:

```go
type Links struct {
  ID        int
  CreatedAt time.Time 
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt 
  ChannelID int
  Link      string
}
```

It contains the unique ID, creation/update timestamps, soft delete flag, channel ID, and link URL for each record.

`GetTestResults` makes concurrent requests to test the latency of each link and returns a slice of `Pair` structs:

```go
type Pair struct {
  ID   int
  Ping int32 
  Link string
}
```

Here is an example:

```go 
var links []models.Links
// populate links from database

results := foreignusage.GetTestResults(&links)

// results contains Pairs with Ping times for each Link
```

The `Ping` field contains the latency result in milliseconds.

This allows efficiently testing the latency of multiple links retrieved from a database concurrently.