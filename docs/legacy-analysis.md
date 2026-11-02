# Comparativo Ponto por Ponto: Código Legado (`museum-code`) vs. Código Atual (Raiz)

Este documento apresenta uma análise detalhada e comparativa entre a versão histórica/legada localizada em [`museum-code/`](file:///d:/ZEDECKLAB/WaysPessoais/feramentas/commit-count/museum-code) (desenvolvida originalmente por Lauripispanen para `commits.top`) e a versão atualizada na raiz do repositório (mantida por Ashkulz para `committers.top`).

---

## 1. Visão Geral das Diferenças Estruturais

A versão legada continha apenas o mecanismo básico de busca e um arquivo de configuração de CI abandonado (`.travis.yml`). A versão atual transformou a ferramenta num **pipeline de dados contínuo (ETL)** totalmente automatizado via GitHub Actions e Cloudflare Workers.

| Aspecto | Versão Legada (`museum-code/`) | Versão Atual (Raiz `/`) | Impacto Arquitetural |
| :--- | :--- | :--- | :--- |
| **Automação CI/CD** | Travis CI (`.travis.yml`) | GitHub Actions (`daily_update.yml`, `deploy_badges.yml`, `dependabot.yml`) | Migração moderna de CI/CD sem custos de serviços externos. |
| **Presets Geográficos** | Apenas `include` e `exclude` (~390 linhas) | Adicionado campo `title`, dezenas de novos países e **Checksum SHA-256** (~642 linhas) | Rastreabilidade e invalidação de cache inteligente. |
| **Integração CLI** | Apenas processamento direto | Nova flag `--list-presets` exportando CSV com checksums | Permite automação externa (scripts Python/Bash). |
| **Formatação YAML** | Timestamp simples `time.Now()` | ISO-8601 (`time.RFC3339`) + metadados de `title` e `definition_checksum` | Compatibilidade com Jekyll/geradores de site estático. |
| **Sistema de Badges** | Não existia | Cloudflare Worker + script de deploy (`badges/`) | Capacidade de servir badges SVG dinâmicos em alta velocidade. |

---

## 2. Comparação Ponto por Ponto dos Arquivos

### Ponto 2.1: `presets.go` (A Evolução da Definição Geográfica)
- **No legado (`museum-code/presets.go`)**:
  - A struct continha apenas:
    ```go
    type QueryPreset struct {
        include []string
        exclude []string
    }
    ```
  - Possuía apenas os primeiros países cadastrados.
- **Na versão atual (`presets.go`)**:
  - A struct ganhou um campo formal de exibição:
    ```go
    type QueryPreset struct {
        title   string
        include []string
        exclude []string
    }
    ```
  - **Função de Checksum SHA-256 (`PresetChecksum`)**:
    Gera um hash SHA-256 combinando o título e todas as cidades incluídas/excluídas do país. Se alguém alterar uma cidade no preset, o checksum muda automaticamente.
  - **Função `PresetTitle`**: Permite recuperar o nome amigável de cada região.

---

### Ponto 2.2: `main.go` (Aprimoramento da CLI para Automação)
- **No legado (`museum-code/main.go`)**:
  - Só aceitava rodar para um preset ou uma lista de localizações arbitrárias.
- **Na versão atual (`main.go`)**:
  - Nova flag CLI:
    ```go
    listPresets := flag.Bool("list-presets", false, "List all available presets as CSV and exit immediately")
    ```
  - Quando disparado com `--list-presets`, ele gera uma saída CSV:
    ```csv
    preset,title,definition_checksum
    angola,"Angola",e3b0c44298fc1c149...
    mozambique,"Mozambique",4a5b6c7d...
    ```
  - Propaga `PresetTitle` e `PresetChecksum` para dentro do pacote `top.Options`.

---

### Ponto 2.3: `top/top.go` e `output/output.go` (Metadados Estruturados)
- **No legado**:
  - O YAML gerava o campo `generated: 2020-05-01 10:00:00...` e apenas `min_followers_required` e `total_user_count`.
- **Na versão atual**:
  - **Padronização RFC3339**: `time.Now().Format(time.RFC3339)`.
  - **Injeção de Metadados**:
    ```go
    if options.PresetTitle != "" && options.PresetChecksum != "" {
        fmt.Fprintf(writer, "title: %+v\n", options.PresetTitle)
        fmt.Fprintf(writer, "definition_checksum: %+v\n", options.PresetChecksum)
    }
    ```
  - Essa injeção permite ao script no GitHub Actions comparar se o arquivo YAML salvo na branch do site tem o mesmo checksum do código Go compilado. Se for diferente, ele sabe que precisa regerar aquela região imediatamente!

---

### Ponto 2.4: A Automação Diária (`.github/workflows/daily_update.yml`)
> [!IMPORTANT]
> Este arquivo **não existe** no código legado (`museum-code/`). Ele é a maior novidade arquitetural do projeto atual.

**O que ele faz:**
1. Compila o binário Go mais recente (`most-active-github-users-counter`).
2. Faz checkout da branch `gh-pages` (onde o website Jekyll reside).
3. Executa um script Python embutido inteligente que:
   - Lê todos os presets via `--list-presets`.
   - Cria as páginas Markdown necessárias (`<país>.md`, `<país>_private.md`, `<país>_public.md` e `rank_only/<país>.json`).
   - Verifica se os arquivos YAML em `_data/locations/*.yml` estão defasados (com mais de 5 dias sem atualização ou com `definition_checksum` desatualizado).
   - Processa em lotes seguros (`MAX_IN_RUN = 3`) para não estourar a cota da API GraphQL do GitHub.
   - Faz commit e push automático direto no repositório.

---

### Ponto 2.5: Sistema de Badges (`badges/`)
Inexistente no legado, o diretório `badges/` contém:
- **`cloudflare_worker.js`**:
  - Um script serverless para Cloudflare Workers.
  - Recebe a requisição pública de imagem (ex: `/users/zedeck/badge.svg`).
  - Consulta o arquivo de ranking JSON gerado pelo GitHub Pages.
  - Devolve um badge em SVG com estilo idêntico ao Shields.io.
- **`deploy`**:
  - Script bash para publicação automatizada do worker na Cloudflare.
- **`.github/workflows/deploy_badges.yml`**:
  - Disparo de deploy automático na Cloudflare sempre que há alteração na pasta `badges/`.

---

## 3. O que Permaneceu Idêntico (Core do Crawler)

Nem tudo mudou. Os seguintes módulos mantiveram a mesma lógica entre as duas versões:
1. **`github/github.go`**:
   - A query GraphQL continua exatamente a mesma.
   - A paginação por `followers:<min` e cursores GraphQL continua idêntica.
   - As métricas de commits, contribuições totais e restritas (privadas) não mudaram.
2. **`net/net.go`**:
   - A composição de middlewares HTTP (`Compose`, `TokenAuth`, `MakeRequester`) permaneceu igual.
3. **`Dockerfile`**:
   - A imagem Docker permaneceu compatível com a compilação do binário Go.

---

## 4. Lições Importantes para a Criação do Seu Clone

Se você vai construir o seu próprio clone, as mudanças entre o legado e o atual deixam claro o caminho:

1. **Evite Atualizar Tudo de Uma Vez**:
   - O legado tentava rodar sem controle de cotas. O projeto atual aprendeu que o limite do GitHub GraphQL (5.000 requisições/hora) exige processamento por lotes (3 a 5 países por execução agendada).
2. **Rastreabilidade por Checksum**:
   - Adicione hashes aos presets para saber quando os termos de pesquisa de uma cidade/país mudaram sem precisar reprocessar o mundo inteiro.
3. **Desacoplamento do Frontend**:
   - O projeto atual gera arquivos de dados puros (`_data/locations/*.yml` e `rank_only/*.json`). Isso permite que você crie um frontend moderno em **Next.js**, **Astro** ou **Vue** consumindo apenas esses dados pré-computados, com custo de servidor praticamente zero.
