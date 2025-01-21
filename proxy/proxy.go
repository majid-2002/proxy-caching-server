package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"caching-proxy/cache"
)

type CacheEntry struct {
	Body    []byte
	Headers http.Header
}

var (
	cacheStorage = cache.NewCache()
	mu           sync.RWMutex
)

func StartServer(port int, origin string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxyHandler(w, r, origin)
	})

	log.Printf("Starting server on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func proxyHandler(w http.ResponseWriter, r *http.Request, origin string) {
	cacheKey := r.URL.String()

	mu.RLock()
	cachedResponse, ok := cacheStorage.Get(cacheKey).(*CacheEntry)
	mu.RUnlock()

	if ok {
		w.Header().Set("X-Cache", "HIT")
		copyHeaders(w, cachedResponse.Headers)
		w.Write(cachedResponse.Body)
		return
	}

	resp, err := http.Get(origin + r.URL.String())
	if err != nil {
		http.Error(w, "Error contacting the origin server", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusBadGateway)
		return
	}

	mu.Lock()
	cacheStorage.Set(cacheKey, &CacheEntry{
		Body:    body,
		Headers: resp.Header.Clone(),
	})
	mu.Unlock()

	// Serve the response
	w.Header().Set("X-Cache", "MISS")
	copyHeaders(w, resp.Header)
	w.Write(body)
}

func copyHeaders(dst http.ResponseWriter, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			dst.Header().Add(key, value)
		}
	}
}

func ClearCache() {
	mu.Lock()
	cacheStorage.Clear()
	mu.Unlock()
}
