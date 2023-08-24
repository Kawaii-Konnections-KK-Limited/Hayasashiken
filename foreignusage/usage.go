package foreignusage

import (
	"sync"

	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/models"
	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/run"
)

type Pair struct {
	ID   int
	Ping int32
	Link string
}

func GetTestResults(links *[]models.LinksSsimplified, t *int32, upperBoundPingLimit *int32, TestUrl *string) []Pair {
	var wg sync.WaitGroup

	var pairs []Pair
	baseBroadcast := ""

	for i, v := range *links {
		link := v.Link
		id := v.ID
		wg.Add(1)
		port := i + 50000
		go func(link *string, port int, i int) {
			defer wg.Done()

			r, _ := run.SingByLink(link, TestUrl, &port, t, &baseBroadcast)
			if r > 10 && r < *upperBoundPingLimit {
				pairs = append(pairs, Pair{
					Ping: r,
					Link: *link,
					ID:   i,
				})
			}

		}(&link, port, id)
	}
	wg.Wait()
	return pairs

}
