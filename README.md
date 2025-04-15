# Template API

A template for building Go APIs with health checks, structured logging, Prometheus metrics, and versioned configuration.

## Features

- **Gin** web framework
- **Prometheus** metrics via middleware
- **Redis** caching
- **Viper** config management
- **Zap** structured logging
- `/healthz` endpoint for health checks

### Routes

**Note:** all custom API routes have the context path `/template-api`.

```
GET    | /healthz
GET    | /metrics

GET    | /AppVersion
```

### Repository file tree

```
├── deploy/
│   ├── config.yaml        # API configuration file
│   ├── deployment.yaml    # Kubernetes manifest file
│
├── internal/
│   ├── api/
│   │   ├── routes.go        # HTTP routes and middleware
│   │
│   ├── config/
│   │   ├── config.go        # Viper configuration logic
│   │
│   ├── controller/
│   │   ├── appVersion.go    # Handler
│   │
│   ├── function/
│   │   ├── appVersion.go    # Business logic
│   │
│   ├── logger/
│   │   ├── logger.go        # Logging configuration
│   │
│   ├── redis/
│   │   ├── client.go        # Redis configuration
│
├── config.yaml    # Local configuration file
├── go.mod         # Go modules file
├── main.go        # Go app entrypoint
│
├── .github/workflows/
│   ├── build.yml
│   ├── lint.yml
│   ├── sast.yml
```
