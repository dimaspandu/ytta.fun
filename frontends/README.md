# Frontends

## Overview

This folder contains **microfrontends** for the YYTA.fun platform. Each app is **domain-oriented** rather than being split by technical layers. All apps are hosted **in isolation** (e.g., subdomains) instead of using module federation.

## Principles

* **Domain-driven apps** → each app focuses on one business domain.
* **Isolated hosting** → apps are deployed independently with separate build pipelines.
* **Shared contracts** → consistency via APIs, shared types, and design-system.
* **Autonomy** → apps can evolve at their own pace as long as API contracts are respected.

## Structure

```
frontends/
  landing-app/       # Public landing page
  blog-app/          # Blog and editorial content
  ads-app/           # Promotions and advertisements
  product-app/       # Product listings with timed publishing
  session-app/       # Group sessions and ordering
  design-system/     # Shared UI components and tokens
```

## Deployment

* Each app is deployed to a dedicated subdomain, e.g.:

  * landing.yyta.fun
  * blog.yyta.fun
  * products.yyta.fun
* Apps do not import each other directly. All communication is through backend APIs.

## Notes

* Use **design-system** for shared UI consistency.
* Each app has its own `README.md` with setup instructions.
