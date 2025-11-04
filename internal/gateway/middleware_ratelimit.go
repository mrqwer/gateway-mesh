// internal/gateway/middleware_ratelimit.go
package gateway

import (
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

var redisClient *redis.Client
var limiter *redis_rate.Limiter

func init() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379" // fallback for local dev
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	limiter = redis_rate.NewLimiter(redisClient)
	log.Printf("✅ Rate limiter initialized with Redis at %s", addr)
}

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		key := r.RemoteAddr

		res, err := limiter.Allow(ctx, key, redis_rate.PerMinute(100))
		if err != nil {
			log.Printf("❌ Redis rate limiter error: %v", err)
			http.Error(w, "rate check error", http.StatusInternalServerError)
			return
		}

		if res.Allowed == 0 {
			w.Header().Set("Retry-After", "60")
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
