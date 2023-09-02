package api

import (
	"context"
	"sync"

	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/run"
)

func getTestResultsAsService(links *[]link, timeout *int32, upperBoundPingLimit *int32, TestUrl *string, ctx *context.Context) []responseLink {
	var wg sync.WaitGroup

	var pairs []responseLink
	baseBroadcast := ""

	for i, v := range *links {
		link := v.Link
		id := v.ID
		kills := make(chan bool, 1)
		wg.Add(1)
		port := i + 50000
		go func(link *string, port int, i int) {
			defer wg.Done()

			r, _ := run.SingByLink(link, TestUrl, &port, timeout, &baseBroadcast, *ctx, &kills)
			if r > 10 && r < *upperBoundPingLimit {
				pairs = append(pairs, responseLink{
					Ping: r,
					Link: *link,
					ID:   i,
				})
			} else {
				kills <- true

			}

		}(&link, port, id)
	}
	wg.Wait()
	return pairs

}
