# Product App

The Product App enables users to create, manage, and browse products. Unlike traditional marketplaces, products here are published for a limited duration (e.g., 24 hours), resembling an open-close store concept.

## Team Scope

* Develop product listing and detail pages.
* Implement the product publishing flow with time-based visibility.
* Ensure a smooth user experience for browsing, filtering, and searching products.
* Optimize for fast loading and responsive design on all devices.

## Tech Stack Recommendations

* **Framework & Build Tools:** Lightweight, fast, and modular for modern frontend development.
* **Styling:** SCSS using BEM (Block Element Modifier) methodology.
* **State Management:** Centralized store for global state handling.
* **Routing:** Modular routing system.
* **Assets:** Optimized images, SVGs, fonts.
* **Build & Performance:** Lazy-loading components, code splitting, minimal JS.

> Note: Focus on modular component design and consistent BEM naming conventions for easy reuse across product pages and future features.

## Suggested Folder Structure

```
product-app/
├─ public/          # Static assets (images, fonts, icons)
├─ src/
│  ├─ assets/       # Images, icons, SVGs, fonts
│  ├─ components/   # Reusable UI components (BEM methodology)
│  ├─ pages/        # Product listing, details, publishing flows
│  ├─ store/        # State management
│  ├─ router/       # Routing configuration
│  ├─ styles/       # Global SCSS
│  └─ utils/        # Helper functions
├─ README.md
└─ package.json
```

## Running the Project

### Install Dependencies

```bash
pnpm install
```

### Run Development Server

```bash
pnpm run dev
```

The app will be available at `http://localhost:5173` (default Vite port).

### Build for Production

```bash
pnpm run build
```

### Preview Production Build

```bash
pnpm run preview
```

The preview server will run on `http://localhost:4173` by default.

## Performance & UX Guidelines

* Minimize JavaScript and use lazy-loading for heavy components.
* Optimize images using modern formats (WebP/AVIF).
* Ensure smooth transitions for product lists and filters.
* Include SEO meta tags for product pages (title, description, OG tags).
* Follow BEM methodology in all component classes for maintainable styling.

## Contribution

* Follow consistent code and component conventions with BEM.
* Use feature branches for updates.
* Test performance and responsiveness before merging.

## License

* [MIT](LICENSE) — Free to use for development, research, and internal projects.
