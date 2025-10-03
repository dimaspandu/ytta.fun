# Landing App

The Landing App is the entry point for new visitors. Its primary goal is to deliver a fast, clear, and engaging introduction to the platform, highlight the core value propositions, and drive conversions such as sign-ups or onboarding.

## Team Scope

* Craft and maintain the public-facing homepage.
* Manage marketing campaigns and seasonal updates on the landing page.
* Ensure the design reflects the product’s identity and overall brand consistency.

## Tech Stack Recommendations

* **Framework:** Astro / Vite + React / SolidStart (lightweight alternatives to Next.js for faster load times)
* **Styling:** SCSS or TailwindCSS
* **Assets:** Optimized images, SVGs, fonts
* **SEO & Performance:** Meta tags, Open Graph, structured data, minimal JS, preloaded critical assets

> Note: For a static landing page, Next.js may introduce unnecessary overhead. Consider lightweight frameworks to improve first paint and Lighthouse scores.

## Suggested Folder Structure

```
landing-app/
├─ public/        # Static assets (images, fonts, icons)
├─ src/
│  ├─ components/  # Reusable UI components
│  ├─ pages/       # Landing pages and routes
│  ├─ styles/      # SCSS/CSS files
│  └─ utils/       # Helper functions
├─ README.md
└─ package.json
```

## SEO & Performance Guidelines

* Minimize JavaScript on landing pages.
* Optimize images using WebP/AVIF formats.
* Preload critical fonts and assets.
* Include meta description, title, and Open Graph tags for social sharing.
* Audit performance with Lighthouse before deployment.

## Contribution

* Follow consistent code and styling conventions.
* Use feature branches for updates or campaign changes.
* Ensure responsive design across devices.

## License

* [MIT](LICENSE) — Free to use for development, research, and internal projects.

---

# Astro Starter Kit: Minimal

```sh
pnpm create astro@latest -- --template minimal
```

> **Seasoned astronaut?** Delete this file. Have fun!

## Project Structure

Inside of your Astro project, you'll see the following folders and files:

```text
/
├── public/
├── src/
│   └── pages/
│       └── index.astro
└── package.json
```

Astro looks for `.astro` or `.md` files in the `src/pages/` directory. Each page is exposed as a route based on its file name.

There's nothing special about `src/components/`, but that's where we like to put any Astro/React/Vue/Svelte/Preact components.

Any static assets, like images, can be placed in the `public/` directory.

## Commands

All commands are run from the root of the project, from a terminal:

| Command                   | Action                                           |
| :------------------------ | :----------------------------------------------- |
| `pnpm install`            | Installs dependencies                            |
| `pnpm dev`                | Starts local dev server at `localhost:4321`      |
| `pnpm build`              | Build your production site to `./dist/`          |
| `pnpm preview`            | Preview your build locally, before deploying     |
| `pnpm astro ...`          | Run CLI commands like `astro add`, `astro check` |
| `pnpm astro -- --help`    | Get help using the Astro CLI                     |

## Want to learn more?

Feel free to check [our documentation](https://docs.astro.build) or jump into our [Discord server](https://astro.build/chat).
