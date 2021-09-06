package httputil

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MethodFilter will filter specified http method
//  allowed is method list that allowed to use
func MethodFilter(allowed []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if len(method) <= 0 {
			method = "GET"
		}
		filtered := true
		for _, allowedMethod := range allowed {
			if method == allowedMethod {
				filtered = false
				break
			}
		}

		if filtered {
			NewError(c, http.StatusMethodNotAllowed, fmt.Errorf("method not allowd: %s", method))
			c.Abort()
		}
	}
}
