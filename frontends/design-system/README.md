# Design System

## Overview

The **design-system** provides shared UI components, styles, and tokens. It ensures **visual and interaction consistency** across all frontend apps.

## Principles

* **Single source of truth** for UI tokens (colors, typography, spacing).
* **Reusable components** for common patterns (buttons, cards, modals).
* **Theming support** for light/dark or brand variations.
* **Framework-agnostic** if possible, but optimized for React (used in microfrontends).

## Structure

```
design-system/
  src/
    components/   # Reusable UI components
    tokens/       # Colors, typography, spacing
    hooks/        # Shared UI logic (if any)
    index.ts      # Entry point
  package.json    # Versioning for internal usage
```

## Usage

* Import components from `@yyta/design-system` (if published as a package).
* Or use relative imports within the monorepo.

Example:

```ts
import { Button } from "@yyta/design-system";

export function Hero() {
  return <Button>Get Started</Button>;
}
```

## Deployment

* Can be published to a private npm registry.
* Alternatively, consumed locally in the monorepo build pipeline.

