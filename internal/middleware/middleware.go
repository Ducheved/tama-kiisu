package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type ipLimiter struct {
	sync.Mutex
	lastAction map[string]map[string]time.Time
}

var limiter = &ipLimiter{
	lastAction: make(map[string]map[string]time.Time),
}

func IPLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			c.Next()
			return
		}

		ip := c.ClientIP()
		action := c.PostForm("action")

		limiter.Lock()
		defer limiter.Unlock()

		if _, exists := limiter.lastAction[ip]; !exists {
			limiter.lastAction[ip] = make(map[string]time.Time)
		}

		lastTime, exists := limiter.lastAction[ip][action]
		now := time.Now()
		if !exists || now.Sub(lastTime) > time.Hour {
			limiter.lastAction[ip][action] = now
			c.Next()
			return
		}

		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded for this action. Try again in an hour."})
		c.Abort()
	}
}
