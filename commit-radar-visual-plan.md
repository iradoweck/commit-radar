# Plan: Frontend Architecture & Visual Flow for Commit Radar

## Goal
Build a clean, neutral, and practical frontend experience where users land on a global country directory, select any country, and instantly view the ranked Developers and Organizations tabs, while preparing the underlying Go engine data contract (JSON per country).

---

## Tasks

- [ ] Task 1: Create the Country Directory Hub on the home view with live search, A-Z index, and neutral country cards. → Verify: Open `http://localhost:2024`, type in search bar, verify matching countries filter dynamically without reload.
- [ ] Task 2: Implement the Country View Router (hash/state-based `#country/name`) with back button, breadcrumbs, and country metadata header. → Verify: Click any country card, URL updates to `#country/{id}`, country header renders with back button.
- [ ] Task 3: Build the Country Tabs component ("Developers Ranking" and "Organizations Ranking"). → Verify: Toggle between "Developers Ranking" and "Organizations Ranking" tabs and check active state classes.
- [ ] Task 4: Design the Top Developers Table (Rank, Avatar, Developer `@login`, Total Contributions, Commits, Pull Requests, Company/Orgs). → Verify: Table renders in Developers tab with clean numeric ranking and hover rows.
- [ ] Task 5: Design the Top Organizations Table (Rank, Organization Name `@org`, Total Associated Devs, Primary Tech/Tags). → Verify: Table renders in Organizations tab sorted from highest member count to lowest.
- [ ] Task 6: Establish the Go-to-Frontend JSON Data Contract (`data/{country}.json` mock & schema definition) so the Go crawler can feed this exact structure without rate-limit issues. → Verify: Verify JSON structure contains `country`, `generated_at`, `total_users`, `users[]`, and `organizations[]`.
- [ ] Task 7: Integrate responsive states and clean empty/loading indicators. → Verify: Resize viewport to mobile/desktop, verify table scroll and card responsiveness.

---

## Done When
- [ ] User opens the site and immediately sees a clean, neutral directory of countries with instant search.
- [ ] Clicking any country opens its dedicated view with clean tabs for "Developers Ranking" and "Organizations Ranking".
- [ ] Ranking displays Developers and Organizations sorted from top to bottom.
- [ ] The data contract is strictly defined for the Go engine to output JSON files per country.

---

## Notes
- **Neutrality:** No hardcoded country bias or pinned flags; 100% alphabetical and search-driven.
- **FinOps & API Quotas:** The Go engine will process in small batches and output static JSON files per country, protecting GitHub token limits.
