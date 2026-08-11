# Gateway docs

Pure API edge for PayLater. Proxies all public routes to domain services.

- Listen address: `PORT` from `gateway/.env`
- Upstreams: `ADMIN_SERVICE_URL`, `MERCHANT_SERVICE_URL`, `CUSTOMER_SERVICE_URL`, `LEDGER_SERVICE_URL`, `REPORT_SERVICE_URL`
- No database configuration
- Build from repo root: `docker build -f gateway/Dockerfile .` (when Dockerfile is added)
