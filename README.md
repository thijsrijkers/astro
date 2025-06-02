# Astro

This project is a low-level implementation of an HTTP/1.1 web server written entirely in Go, without using the standard `net/http` package for core functionality. It handles raw TCP connections, manually parses HTTP requests, routes incoming requests to handlers, and constructs valid HTTP responses. The server supports persistent connections (keep-alive), concurrency via goroutines, and proper HTTP/1.1 compliance. Built from first principles, this server is designed to be minimal, educational, and extensible—ideal for learning the internals of web protocols or building lightweight, custom server software.

## Goals

- Learn and implement the internal workings of a web server
- Build a modular, extensible core for HTTP/1.1
- Prepare for production-level features through clean, low-level implementation

---

## Phase 1: Minimal, Functional HTTP/1.1 Server

This phase includes everything needed to create a working web server that can accept, parse, and respond to HTTP/1.1 requests from clients like browsers or `curl`.

### Core Components

1. **TCP Listener**
   - Open a TCP socket on a given port and accept client connections.

2. **Connection Handling**
   - Handle each connection in its own goroutine for concurrency.
   - Support multiple simultaneous client connections.

3. **HTTP Request Parsing**
   - Parse the HTTP request line, headers, and optional body.
   - Store method, path, version, headers, and body in structured form.

4. **Basic Routing**
   - Match HTTP method and path.
   - Route requests to appropriate handler functions.

5. **HTTP Response Construction**
   - Manually construct valid HTTP/1.1 responses including status line, headers, and body.
   - Set `Content-Length`, `Content-Type`, and other required headers.

6. **Keep-Alive Support**
   - Respect `Connection: keep-alive` and `Connection: close` headers.
   - Allow persistent connections as defined by HTTP/1.1.

7. **Error Handling**
   - Return appropriate error responses (400, 404, 500).
   - Log malformed requests and server errors.

8. **Graceful Shutdown**
   - Catch termination signals.
   - Close open connections and listener safely.

---

## Phase 2: Full-Featured HTTP Server (Optional Enhancements)

After completing the minimal server, the following steps can be taken to build a production-ready feature set.

### Protocol Support

- Add support for HTTP/2
- Implement WebSocket upgrade handling

### Static and Dynamic Content

- Serve static files with proper MIME type detection
- Implement support for file uploads
- Add templating support for dynamic responses

### Security Features

- Add TLS (HTTPS) support using self-signed or real certificates
- Implement rate limiting
- Add IP filtering and user-agent filtering
- Set secure headers like HSTS and CSP

### HTTP Features

- Support for chunked transfer encoding
- Implement gzip compression for responses
- Add cookie parsing and setting
- Handle redirects (301, 302)

### Middleware and Routing

- Build a middleware stack for logging, authentication, and metrics
- Create a routing layer with parameterized routes

### Observability

- Add structured logging
- Collect basic metrics (requests per second, latency)
- Support request tracing with correlation IDs

### Configuration and Deployment

- Use configuration files or environment variables for settings
- Add support for graceful restarts or hot reload during development
- Dockerize the application for container deployment

---

## License

This project is licensed under the MIT License.
