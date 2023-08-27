package foreignusage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"sync"
	"syscall"

	"os"

	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/run"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitService() {

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println(errors.New("please set PORT environment variable"))
		port = "8080"
		fmt.Println("Defaulting to port ", port)

	}
	err := initRouter().Run(":" + port)
	if err != nil {
		return
	}
}
func initRouter() *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r := gin.New()

	r.Use(cors.New(config))
	// r.GET("/", api.RestrictedEndpoint)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/test", testHandler)

	return r
}

type link struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}
type requestLinks struct {
	Links               *[]link `json:"links"`
	Timeout             *int32  `json:"timeout"`
	UpperBoundPingLimit *int32  `json:"upperbound"`
	TestUrl             *string `json:"testurl"`
}
type responseLink struct {
	ID   int    `json:"id"`
	Ping int32  `json:"ping"`
	Link string `json:"link"`
}
type responseLinks struct {
	Links []responseLink `json:"links"`
}
type responseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func testHandler(c *gin.Context) {
	var r requestLinks
	err := c.BindJSON(r)
	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	go func() {

		<-sigCh

		fmt.Println("Received sigint signal")

		cancel() // Cancel the context when SIGINT is received

	}()
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{Status: http.StatusBadRequest, Message: "Invalid request."})
		return
	}
	if c.GetHeader("Authorization") == os.Getenv("auth") {
		res := getTestResultsAsService(r.Links, r.Timeout, r.UpperBoundPingLimit, r.TestUrl, &ctx)
		c.JSON(http.StatusBadRequest, responseLinks{Links: res})
	} else {
		c.JSON(http.StatusBadRequest, responseError{Status: http.StatusForbidden, Message: "request unauthorized."})

	}

}
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
