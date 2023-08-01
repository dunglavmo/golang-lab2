// package loadbalancer

// import (
// 	"github.com/spf13/viper"
// )

// var backendServers []string
// var rateLimit int
// var loggingEnabled bool

// func initConfig() error {
// 	viper.SetConfigFile("config/system.conf")
// 	if err := viper.ReadInConfig(); err != nil {
// 		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
// 			// Config file not found; ignore error if desired
// 		} else {
// 			// Config file was found but another error was produced
// 		}
// 	}

// 	backendServers = viper.GetStringSlice("backend_servers")
// 	rateLimit = viper.GetInt("rate_limit")
// 	loggingEnabled = viper.GetBool("logging_enabled")

//		return nil
//	}
//
// loadbalancer/config.go
package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/time/rate"
)

type Config struct {
	BackendServers []string
	RateLimit      float64
	LoggingEnabled bool
	Port           int
}
type LoadBalancer struct {
	Config  *Config
	Limiter *rate.Limiter
}

func ReadConfig() (*Config, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config/system.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{
		BackendServers: viper.GetStringSlice("backend_servers"),
		RateLimit:      viper.GetFloat64("rate_limit"),
		LoggingEnabled: viper.GetBool("logging_enabled"),
		Port:           viper.GetInt("port"),
	}

	return config, nil
}
func NewLoadBalancer(config *Config) *LoadBalancer {
	return &LoadBalancer{
		Config:  config,
		Limiter: rate.NewLimiter(rate.Limit(config.RateLimit), 1),
	}
}

func (lb *LoadBalancer) Handler(w http.ResponseWriter, r *http.Request) {
	// Check rate limit for the incoming request's IP address
	if lb.Limiter.AllowN(time.Now(), 1) == false {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	// Round Robin to select the next backend server
	backendURL, err := url.Parse(lb.Config.BackendServers[0])
	if err != nil {
		http.Error(w, "Error parsing backend URL", http.StatusInternalServerError)
		return
	}

	lb.Config.BackendServers = append(lb.Config.BackendServers[1:], lb.Config.BackendServers[0])

	// Log the request if logging is enabled
	if lb.Config.LoggingEnabled {
		log.Printf("Load Balancer: Received request from %s for URL: %s\n", r.RemoteAddr, r.URL.Host)
	}

	// Proxy the request to the selected backend server
	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	proxy.ServeHTTP(w, r)

	// Log the response if logging is enabled
	if lb.Config.LoggingEnabled {
		log.Printf("Load Balancer: Responded to %s for URL: %s\n", r.RemoteAddr, r.URL)
	}
}
