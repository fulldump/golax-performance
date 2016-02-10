package main

import (
	"log"
	"strconv"
	"time"

	"model"

	"github.com/fulldump/golax"
)

/**
 * Interceptor {user_id}
 * if: `user_id` exists -> load the object and put it available in the context
 * else: raise 404
 */
var interceptor_user = &golax.Interceptor{
	Before: func(c *golax.Context) {
		user_id, _ := strconv.Atoi(c.Parameter)
		if user, exists := model.Users[user_id]; exists {
			c.Set("user", user)
		} else {
			c.Error(404, "user `"+c.Parameter+"` does not exist")
		}
	},
}

/**
 * Interceptor logger
 * Stores the time before and log a line for each request:
 * 2016/02/10 23:59:42 GET	/service/v1/users	33.277µs
 * 2016/02/10 23:59:42 GET	/service/v1/users/2	7.483µs
 */
var interceptor_logger = &golax.Interceptor{
	Before: func(c *golax.Context) {
		c.Set("logger_start", time.Now())
	},
	After: func(c *golax.Context) {
		logger_start, _ := c.Get("logger_start")
		log.Printf(
			"%s\t%s\t%s",
			c.Request.Method,
			c.Request.RequestURI,
			time.Since(logger_start.(time.Time)),
		)
	},
}
