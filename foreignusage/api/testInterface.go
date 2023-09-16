package api

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/pkg/cache"
	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/run"
)

func getTestResultsAsService(links *[]link, timeout *int32, upperBoundPingLimit *int32, TestUrl *string, ctx *context.Context) []responseLink {
	var wg sync.WaitGroup

	var pairs []responseLink
	baseBroadcast := ""
	ca := cache.GetCache()
	for i, v := range *links {
		link := v.Link
		id := v.ID
		if val, exists := ca.Get(link); !exists {

			kills := make(chan bool, 1)
			wg.Add(1)
			port := i + 50000
			go func(link *string, port int, i int) {
				defer wg.Done()

				r, err := run.SingByLink(link, TestUrl, &port, timeout, &baseBroadcast, *ctx, &kills)
				ca.Set(*link, r, 5*time.Minute)
				if r > 10 && r < *upperBoundPingLimit {
					pairs = append(pairs, responseLink{
						Ping: r,
						Link: *link,
						ID:   i,
					})
				} else {
					kills <- true

				}
				if err != nil {
					fmt.Println(err)
				}

			}(&link, port, id)
		} else {
			// fmt.Println(exists)
			if val.(int32) > 10 && val.(int32) < *upperBoundPingLimit {

				pairs = append(pairs, responseLink{
					Ping: val.(int32),
					Link: link,
					ID:   i,
				})
			}
		}
	}
	wg.Wait()
	return pairs

}
