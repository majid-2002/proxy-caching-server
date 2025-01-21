### **Caching Proxy Server**  

#### **Project Description**  
This project is a **Caching Proxy Server**, designed to act as an intermediary between a client and an origin server. The server forwards client requests to the origin server, caches the responses, and serves subsequent requests from the cache if available. By implementing caching, this proxy reduces server load, minimizes latency, and improves overall efficiency.

---

#### **How It Works**  

1. **Proxy Functionality**:  
   - A **proxy server** is an intermediary that relays client requests to the intended destination (origin server). It acts as a gateway, forwarding requests and returning responses to the client.
   - In this project, the proxy server listens on a specified port, processes incoming HTTP requests, and forwards them to the configured origin server if the response is not already cached.

2. **Caching Mechanism**:  
   - Caching temporarily stores server responses so that identical requests in the future can be served directly from the cache rather than re-fetching the data from the origin server.
   - This caching proxy implements in-memory caching using a map-like structure, associating URLs (as keys) with the corresponding responses (as values).

3. **Response Workflow**:
   - **On Request**:
     - When a client makes a request, the proxy checks if the response is already cached.
     - If the response exists in the cache:
       - It is returned immediately to the client with the header `X-Cache: HIT`.
     - If the response is not cached:
       - The proxy forwards the request to the origin server.
       - The origin server's response is cached and then returned to the client with the header `X-Cache: MISS`.
   - Cached responses include the HTTP body and headers to replicate the origin server’s response accurately.

4. **Cache Clearing**:  
   - A command-line flag `--clear-cache` is implemented to clear all cached entries, allowing for cache management without restarting the server.

---


#### **Key Concepts**

1. **Proxy Server**:  
   - A proxy server acts as a middleman between a client and an origin server. It can be used for security, performance, or load balancing purposes.  
   - In this project, the proxy also handles caching, reducing redundant requests to the origin server.

2. **Caching**:  
   - Caching is the process of storing copies of files or data in a temporary storage layer (cache) so future requests can be served faster.
   - It improves performance by reducing latency (time taken to fetch data) and saving bandwidth between the client and origin server.

---

#### **How It Works - Simplified Workflow**

1. **Request Handling**:
   - The client sends an HTTP request to the caching proxy server.
   - The proxy checks if the request is in its cache.

2. **Cache Hit**:
   - If the requested data exists in the cache, it is returned immediately with the header `X-Cache: HIT`.

3. **Cache Miss**:
   - If the requested data is not in the cache:
     - The proxy forwards the request to the origin server.
     - The origin server processes the request and sends a response back to the proxy.
     - The proxy caches the response (body and headers) for future use.
     - The proxy sends the response to the client with the header `X-Cache: MISS`.

4. **Clearing the Cache**:
   - Use the `--clear-cache` flag while running the server to purge all cached entries.

---

#### **Potential Use Cases**  
- **Content Delivery Optimization**:
  - Cache frequently accessed data (e.g., API responses, static resources) to improve performance.
- **Reduced Load on Origin Servers**:
  - Minimize redundant requests to the backend server by serving data directly from the cache.
- **Latency Reduction**:
  - Serve cached responses faster than fetching them from the origin server.
