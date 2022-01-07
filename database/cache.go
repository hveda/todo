package database

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"

	memcache "github.com/patrickmn/go-cache"
)

var TCache = memcache.New(5*time.Second, 10*time.Second)
var MemCache = memcache.New(10*time.Second, 15*time.Second)

// Cache will return a caching middleware
func Cache(exp time.Duration) fiber.Handler {
	return cache.New(cache.Config{
		Expiration:   exp,
		CacheControl: true,
	})
}
