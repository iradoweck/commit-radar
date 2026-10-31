# Museu & Livro de Memória: Créditos aos Criadores e Mantenedores

Este repositório homenageia e preserva a linhagem histórica dos projetos que inspiraram a nossa **Versão 3 (Engenharia Reversa)**.

---

## 🏛️ 1. O Criador Original: Lauri Piispanen (`lauripiispanen`)
- **Diretório Histórico:** [`museum-code/lauripiispanen-v1-original/`](file:///d:/ZEDECKLAB/WaysPessoais/feramentas/commit-count/museum-code/lauripiispanen-v1-original)
- **Repositório Inicial:** `lauripiispanen/most-active-github-users-counter` (utilizado para alimentar o domínio pioneiro `commits.top`).
- **Contribuição Fundamental:**
  - Concepção da ideia de contornar a limitação da API do GitHub (que não permite ordenar diretamente por commits) ordenando preliminarmente por seguidores decrescentes (`sort:followers-desc`) e extraindo o calendário de contribuições via GraphQL.
  - Implementação dos algoritmos em Go de classificação e agregação de organizações.

---

## 🚀 2. O Mantenedor da Evolução: Ashish Kulkarni (`ashkulz`)
- **Diretório Histórico:** [`museum-code/ashkulz-v2-upstream/`](file:///d:/ZEDECKLAB/WaysPessoais/feramentas/commit-count/museum-code/ashkulz-v2-upstream)
- **Repositório Upstream:** `ashkulz/committers.top` (engine que manteve o projeto ativo no ar via `committers.top`).
- **Contribuição Fundamental:**
  - Transformação do script manual em um pipeline contínuo de dados (ETL) no GitHub Actions (`daily_update.yml`).
  - Criação do sistema de integridade de presets com Checksum SHA-256 (`definition_checksum`).
  - Implementação da infraestrutura de Badges dinâmicos no Edge com Cloudflare Workers.

---

## ⚡ 3. A Nova Era: Nossa Versão (Engenharia Reversa na Raiz `/`)
- **Objetivo:**
  - Redesenhar a arquitetura para suportar um ecossistema moderno.
  - Foco especial em granularidade regional refinada (África Lusófona, Moçambique, Angola, Brasil e regiões globais).
  - Interface rica com histórico temporal, persistência otimizada e dashboards modernos.
