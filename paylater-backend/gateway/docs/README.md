# Gateway docs

API gateway for PayLater Phase 4. Proxies admin and merchant microservices and hosts remaining customer, ledger, and report routes until those services are extracted.

- Default listen address: `:9090`
- Upstream: `ADMIN_SERVICE_URL`, `MERCHANT_SERVICE_URL`
- Build from repo root: `docker build -f gateway/Dockerfile .`
