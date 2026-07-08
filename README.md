# Bible Go API

[![Tests](https://github.com/rkeplin/bible-go-api/actions/workflows/test.yml/badge.svg)](https://github.com/rkeplin/bible-go-api/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/rkeplin/bible-go-api/branch/master/graph/badge.svg)](https://codecov.io/gh/rkeplin/bible-go-api)

Bible Go API is an open source REST API containing multiple translations of The Holy Bible, as well as cross-references.
This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY.

### Live Demo

A live demo can be viewed at [https://bible-go-api.rkeplin.com/v1/books/1/chapters/1](https://bible-go-api.rkeplin.com/v1/books/1/chapters/1).

---

## Getting Started (Local Development)

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose v2](https://docs.docker.com/compose/)
- [Go 1.24+](https://go.dev/dl/) (only needed to run tests locally without Docker)

### Run with Docker Compose

```bash
git clone https://github.com/rkeplin/bible-go-api
cd bible-go-api
make dev
```

The API will be available at [http://localhost:8084](http://localhost:8084).

> Note: On first start the MariaDB volume may take several seconds to initialize before queries succeed.

### Hot Reload (Dev Mode)

```bash
make watch
```

### Stop

```bash
make down
```

---

## Running Tests

```bash
make test
```

This runs tests inside the dev Docker image (Go 1.24), so no local Go installation is required. The stack does not need to be running — the test container starts and removes itself automatically.

Tests do not require a running database or Elasticsearch — they run entirely in-process using mock dependencies.

---

## Building

### Build the Docker image

```bash
make build
```

### Build the binary directly

```bash
CGO_ENABLED=0 go build -o server .
```

---

## Deploying to Kubernetes

### Prerequisites

- A running Kubernetes cluster with `kubectl` configured
- [cert-manager](https://cert-manager.io/) installed (for TLS)
- [ingress-nginx](https://kubernetes.github.io/ingress-nginx/) installed

### 1. Configure secrets

Copy `.env.example` to `.env` and fill in real values:

```bash
cp .env.example .env
# edit .env with your credentials
```

### 2. Deploy

```bash
make k8s-deploy
```

This runs:
- `kubectl apply` for the `bible` namespace
- Creates a `bible-env` secret from your `.env` file
- Deploys MariaDB, Elasticsearch, the API, and the Ingress

### 3. Check status

```bash
make k8s-status
```

### 4. Tear down

```bash
make k8s-delete
```

> **Note:** `.env` is gitignored. Never commit it.

---

## API Reference

### Translations

```
GET /translations
GET /translations/{id}
```

### Genres

```
GET /genres
GET /genres/{id}
```

### Books & Chapters

```
GET /books
GET /books/{id}
GET /books/{id}/chapters
GET /books/{bookId}/chapters/{chapterId}
GET /books/{bookId}/chapters/{chapterId}/{verseId}
```

Supply `?translation=KJV` (or any abbreviation) as a query parameter to change the translation. Default is KJV.

Supported translations: `ASV`, `BBE`, `DBY`, `KJV`, `WEB`, `YLT`, `ESV`, `NIV`, `NLT`, `NLT2015`

### Cross References

```
GET /verse/{verseId}/relations
```

### Search

```
GET /search?query={query}[&translation=KJV&offset=0&limit=100]
GET /searchAggregator?query={query}[&translation=KJV]
```

---

## Related Projects

- [Bible Go API](https://github.com/rkeplin/bible-go-api)
- [Bible UI (React)](https://github.com/rkeplin/bible-ui)
- [Bible MariaDB Docker Image](https://github.com/rkeplin/bible-mariadb)

## Credits

Data sourced from:
- [scrollmapper/bible_databases](https://github.com/scrollmapper/bible_databases)
- [honza/bibles](https://github.com/honza/bibles)
