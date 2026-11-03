# Changelog

All notable changes and milestones in the **Most Active GitHub Users (Commit Rank)** project are documented here by date.

---

### 2026-11-03 — Version 3.0.1

- **Interactive Metric Sorting & Dynamic Re-Ranking**
  - Added metric filter selector: `All (Total)`, `Commits`, and `Contributions`.
  - Column headers in the developers table are now interactive and clickable with indicator carets (`▼`).
  - Implemented dynamic rank recalculation: ranks (`#1, #2, #3...`) dynamically re-order based on the selected metric.
  - Multi-page pagination (100 users per page up to Top 300) fully supports and preserves active sort order.

- **Organization Enhancements**
  - Expanded organization extraction capacity to Top 50 organizations.
  - Organization ranking table now fetches and displays real organization logos directly from GitHub with direct profile links.

- **Crawler & Query Throughput Optimization**
  - Increased GraphQL query page size to 15 users and reduced sub-node organizations to 10, eliminating `RESOURCE_LIMITS_EXCEEDED` errors and cutting crawler execution time down by over 60%.
  - Fixed natural pagination termination in Go CLI to avoid fake retries on completed searches.

- **Repository Hygiene & Legacy Code Cleanup**
  - Completely removed archived `museum-code/` folder to eliminate duplicate Go manifests, Dockerfiles, and stale workflows.
  - Historical attributions and memorial links for [@lauripiispanen](https://github.com/lauripiispanen) and [@ashkulz](https://github.com/ashkulz) are permanently preserved in the UI and README.
  - Cleaned `.gitignore` and `.geminiignore` to optimize build and token footprint.

---

### 2026-11-01 — Version 3.0.0

- **Dedicated Documentation Page (`#documentation`)**
  - Built a comprehensive integration guide covering profile and repository setup.
  - Added live dynamic Shields.io badge specifications with automatic unranked fallback.
  - Published the auditable GitHub GraphQL indexing algorithm and filter criteria.

- **Open Source Technologies Showcase**
  - Redesigned the technologies modal into a clean, responsive 2-column layout.
  - Highlighted full underlying stack: Go, GitHub GraphQL API, Tailwind CSS, Python, Inter, JetBrains Mono, Shields.io, GitHub Actions, CodeQL, and Ko-fi SDK.

- **Community Support & Infrastructure Goal**
  - Integrated dedicated support modal outlining the custom domain and telemetry server infrastructure goal.
  - Configured seamless Ko-fi checkout flow allowing direct contributions without leaving the page.
  - Added maintainer profile badge with avatar and GitHub Sponsors alternative.

- **Header Live Metrics & Navigation**
  - Added live GitHub PR and Issue counters fetched dynamically from the official GitHub API.
  - Introduced clean top navigation with direct access to About, Documentation, and Repository source.

---

### 2026-10-31

- **Country Directory & Drill-Down Views**
  - Built high-speed client-side search across worldwide countries with an A-Z quick jump bar.
  - Implemented dual ranking views: **Developers Ranking** (contributions, commits, PRs, company) and **Organizations Ranking** (associated developers, impact ratio).
  - Added regional collapsible accordion with localized dynamic badges, pseudocode extraction logic, and single-click JSON datasets.

- **DevSecOps, CI/CD & Repository Security**
  - Configured GitHub Actions CI pipeline running Go build suites on `main` and `beta`.
  - Implemented GitHub CodeQL static analysis workflows for continuous vulnerability detection.
  - Configured Dependabot automated versioning and security updates.
  - Published comprehensive `SECURITY.md` vulnerability reporting guidelines.
  - Configured `.geminiignore` adhering to FinOps principles to optimize AI token context and prevent payload waste.
