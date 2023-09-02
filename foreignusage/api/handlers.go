package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func testHandler(c *gin.Context) {

	var r requestLinks
	err := c.BindJSON(&r)
	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan bool)
	go func() {

		select {
		case <-done:
			fmt.Println("test done")
			cancel()

		}

	}()

	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{Status: http.StatusBadRequest, Message: fmt.Sprint(err)})
		return
	}

	if c.GetHeader("Authorization") == os.Getenv("auth") {
		res := getTestResultsAsService(&r.Links, &r.Timeout, &r.UpperBoundPingLimit, &r.TestUrl, &ctx)
		c.JSON(http.StatusOK, responseLinks{Links: res})
		done <- true

	} else {
		c.JSON(http.StatusBadRequest, responseError{Status: http.StatusForbidden, Message: "request unauthorized."})

	}

}

func pingSelf(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
