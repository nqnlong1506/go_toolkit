# GO_TOOLKIT

## microservices
```
.
├── api-gateway
│   ├── cmd
│   ├── configs
│   └── internal
│       ├── aggregator
│       ├── client
│       ├── handler
│       └── middleware
├── deployments
├── go.mod
├── pkg
│   ├── interceptor
│   ├── logger
│   └── response
├──proto
│   ├── customer
│   └── order
├── services
│   ├── customer
│   │   ├── cmd
│   │   ├── configs
│   │   └── internal
│   │       ├── app
│   │       ├── domain
│   │       ├── infrastructure
│   │       │   ├── grpc_client
│   │       │   └── repository
│   │       └── transport
│   │           ├── grpc
│   │           └── http
│   └── order
│       ├── cmd
│       ├── configs
│       └── internal
│           ├── app
│           ├── domain
│           ├── infrastructure
│           │   ├── grpc_client
│           │   └── repository
│           └── transport
│               ├── grpc
│               └── http
└── toolkit
```
41 directories, 2 files
