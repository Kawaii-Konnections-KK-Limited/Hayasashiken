# Hayasashiken
simple speed test

# foreignusage

This package provides functions for testing network latency of links stored in a database. 

## Usage

```go
import "github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/foreignusage"
```

Here is the function documentation rewritten in Markdown:

## `GetTestResults`

Performs latency tests on a list of links and returns the results.

### Parameters

- `links` - A slice of `LinksSimplified` structs containing the links to test
- `t` - The timeout in milliseconds for the ping requests  
- `upperBoundPingLimit` - The upper bound ping limit in milliseconds
- `TestUrl` - The base URL to ping for the latency tests

### Returns

A slice of `Pair` structs containing the results, where each `Pair` contains:

- `Ping` - The ping latency result 
- `Link` - The tested link
- `ID` - The ID of the link

### Example

```go
links := []LinksSimplified{
  {ID: 1, Link: "http://example.com"},
  {ID: 2, Link: "http://test.com"},
}

results := GetTestResults(links, 50, 100, "http://ping.com")

// results contains latency measurements for each link
```

### Description

The function concurrently pings each link using a goroutine. It waits for all ping requests to complete before returning the aggregated results. Any ping latency over 10ms and under the upper bound limit is captured as a `Pair` in the result slice.

It contains the unique ID, creation/update timestamps, soft delete flag, channel ID, and link URL for each record.

`GetTestResults` makes concurrent requests to test the latency of each link and returns a slice of `Pair` structs:

```go
type Pair struct {
  ID   int
  Ping int32 
  Link string
}
```

The `Ping` field contains the latency result in milliseconds.

This allows efficiently testing the latency of multiple links retrieved from a database concurrently.