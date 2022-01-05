package cutoff

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestFindMatchPath(t *testing.T) {
	//gin.SetMode(gin.ReleaseMode)
	
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users/:userID", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	pathRegexps := GenRouteRegexp(r)
	uri := "/users/234234"
	result := FindMatchPath("GET", uri, pathRegexps)
	target := "/users/:userID"
	if target == result {
		t.Logf("success, %v", target)
	} else {
		t.Errorf("error")
	}
}
