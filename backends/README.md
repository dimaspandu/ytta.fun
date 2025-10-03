# Backends

## Overview

This folder contains **microservices** that power YYTA.fun. Each service is **self-contained**, managing its own data and business rules.

## Principles

* **Single responsibility** → one service = one bounded context.
* **Loose coupling** → services communicate via APIs (REST/gRPC) or events.
* **Resilience** → failure in one service should not bring down the entire system.
* **Consistency via contracts** → shared schemas and API gateway.

## Structure

```
backends/
  user-service/      # User profiles and accounts
  product-service/   # Product lifecycle
  session-service/   # Session lifecycle
  order-service/     # Order management
  payment-service/   # Payments and QR handling
  ads-service/       # Ads campaign data
  blog-service/      # Blog and editorial management
```

## Communication

* Services expose REST APIs with standardized error formats.
* Potential future use of message broker for async flows.
* An API Gateway aggregates and secures public endpoints.

## Deployment

* Each service can be scaled independently.
* Healthcheck endpoints available at `/health`.
* Logs should follow a unified format for observability.