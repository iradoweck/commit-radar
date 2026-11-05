# Commit-Rank — Most Active GitHub Users Counter

This CLI tool and web platform queries the GitHub GraphQL API for developers and organizations, ranking them transparently according to their public contributions, commits, pull requests, and followers across global and regional presets.

---

## 🌍 Open Source Purpose & Global Vision

**Commit-Rank** is maintained as an open, neutral, and global analytics tool under the **Zedecks Open Source** initiative, building upon the foundational work originally created by [@lauripiispanen](https://github.com/lauripiispanen/most-active-github-users-counter) and maintained by [@ashkulz](https://github.com/ashkulz/committers.top).

Our core purpose is to:
1. **Provide Neutral Global Visibility:** Highlight open-source consistency and engineering contributions across every continent without geographical bias.
2. **Preserve Upstream Compatibility:** Maintain the battle-tested Go CLI and GraphQL extraction architecture while modernizing the runtime (Go 1.24+), HTTP context timeouts, segmented location deduplication, and automated JSON dataset pipelines.
3. **Deliver a Fast, Dependency-Free Web Experience:** Offer a clean, mobile-first web interface (`index.html`) backed by static JSON datasets updated automatically via GitHub Actions.

---

## 🚀 Usage & CLI Engine

### GitHub Token

To make requests against the GitHub GraphQL API, you need a personal access token, which can be created [here](https://github.com/settings/tokens). The token requires `read:org` and `read:user` permissions.

### Example Usage (Development Environment)

```bash
go run *.go \
   --token paste-your-token-here \
   --preset worldwide \
   --amount 100 \
   --consider 300 \
   --orgs 100 \
   --output json \
   --file ./data/countries/worldwide.json
```

Supported `--output` formats: `plain`, `csv`, `yaml`, `json`.

---

## 🔄 Automated Ranking Schedule

To respect GitHub API rate limits while keeping datasets fresh, [`.github/workflows/update-rankings.yml`](.github/workflows/update-rankings.yml) runs a rotating daily schedule:
- **Long Presets (High-volume regions):** Dedicated execution days for `worldwide`, `united-states`, and `brazil`.
- **Short Presets (Regional batches):** Grouped into batches of 4–5 countries per run across Africa (`mozambique`, `angola`, `nigeria`, `south-africa`, `kenya`), North America & Caribbean (`canada`, `mexico`, `the-bahamas`, `panama`), and South America (`argentina`, `colombia`, `venezuela`, `uruguay`).
- **Manual Trigger:** Any specific preset (or `all`) can be triggered on-demand via `workflow_dispatch`.

---

## 🤝 Contribution

Contributions are welcome! Please report issues or open pull requests against `main`.

*Please ensure your code is formatted and passes static checks (`go fmt ./...` and `go vet ./...`) before submitting a PR.*

### Preset Changes

When sending PRs for changes to `presets.go`:
1. Ensure all strings and comments use strict UTF-8 encoding.
2. Test or trigger the [Update Country Rankings workflow](.github/workflows/update-rankings.yml) via `workflow_dispatch` with your target preset name.
3. Reference the generated dataset changes in your pull request description.

---

## ❓ FAQ

### Why is this project maintained here?

The [original project](https://github.com/lauripiispanen/most-active-github-users-counter) by [@lauripiispanen](https://github.com/lauripiispanen) was archived, and [@ashkulz](https://github.com/ashkulz/committers.top) kept it alive for the community via `committers.top`. This repository continues that open-source lineage with a modernized Go 1.24 engine, segmented location search (preventing multi-location `AND` conjunction zero-result bugs), and an integrated interactive web directory supported by **Zedecks Open Source**.

### Why am I not on this list?

This could be due to a few factors:

1. **Follower Threshold:** The GitHub Search API does not allow sorting directly by contributions. Instead, the tool first queries users sorted by number of followers (`--consider`) and then ranks that sample by contributions (`--amount`). Each country dataset records the `min_followers_required` threshold.
2. **Location String Matching:** GitHub location search is free-text and not strictly geographical. If your city is not listed in `presets.go`, adding your country name or state to your GitHub profile location will ensure you are matched.
3. **Public vs. Private Contributions:** By default, the main ranking prioritizes public open-source contributions (`PublicContributionCount`).

### Why is my contribution count different from my GitHub profile?

Depending on your profile settings, GitHub displays either only public contributions or both public and private (`restrictedContributionsCount`) contributions. You can inspect what the GitHub GraphQL API returns for your account using the [GitHub GraphQL Explorer](https://docs.github.com/en/graphql/overview/explorer):

```graphql
query { 
  viewer { 
    login
    contributionsCollection {
      restrictedContributionsCount
      contributionCalendar {
        totalContributions
      }
    }
  }
}
```
