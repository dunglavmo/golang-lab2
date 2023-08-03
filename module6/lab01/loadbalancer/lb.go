package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"loadbalancer/repositories"

	"github.com/spf13/viper"
	"golang.org/x/time/rate"
)

type Config struct {
	BackendServers []string
	RateLimit      int
	LoggingEnabled bool
	Port           int
}
type LoadBalancer struct {
	Config  *Config
	Limiter *rate.Limiter
	repo    *repositories.RedisRepository
}

func ReadConfig() (*Config, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config/system.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{
		BackendServers: viper.GetStringSlice("backend_servers"),
		RateLimit:      viper.GetInt("rate_limit"),
		LoggingEnabled: viper.GetBool("logging_enabled"),
		Port:           viper.GetInt("port"),
	}

	return config, nil
}
func NewLoadBalancer(config *Config, repo *repositories.RedisRepository) *LoadBalancer {
	return &LoadBalancer{
		Config:  config,
		Limiter: rate.NewLimiter(rate.Limit(config.RateLimit), 1),
		repo:    repo,
	}
}

func (lb *LoadBalancer) Handler(w http.ResponseWriter, r *http.Request) {

	ip := r.RemoteAddr
	if lb.repo.IsRateLimited(ip, lb.Config.RateLimit) {
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
		log.Printf("Load Balancer: Received request from %s for URL: %s\n", ip, r.URL.Host)
		log.Println(r.RemoteAddr)
	}

	// Proxy the request to the selected backend server
	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	proxy.ServeHTTP(w, r)

	// Log the response if logging is enabled
	if lb.Config.LoggingEnabled {
		log.Printf("Load Balancer: Responded to %s for URL: %s\n", ip, r.URL)
	}
}
