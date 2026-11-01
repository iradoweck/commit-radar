# Changelog

All notable changes and milestones in the **Most Active GitHub Users** project are documented here by date.

---

### 2026-11-01

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

- **Heritage Preservation & Attributions**
  - Archived foundational V1 implementation by [@lauripiispanen](https://github.com/lauripiispanen) in `museum-code/lauripiispanen-v1-original/`.
  - Archived upstream V2 implementation by [@ashkulz](https://github.com/ashkulz) in `museum-code/ashkulz-v2-upstream/`.
  - Added memorial credits (`museum-code/CREDITS.md`) and historical context documenting the evolution into V3 maintained by [@iradoweck](https://github.com/iradoweck).

- **DevSecOps, CI/CD & Repository Security**
  - Configured GitHub Actions CI pipeline running Go build suites on `main` and `beta`.
  - Implemented GitHub CodeQL static analysis workflows for continuous vulnerability detection.
  - Configured Dependabot automated versioning and security updates.
  - Published comprehensive `SECURITY.md` vulnerability reporting guidelines.
  - Configured `.geminiignore` adhering to FinOps principles to optimize AI token context and prevent payload waste.
