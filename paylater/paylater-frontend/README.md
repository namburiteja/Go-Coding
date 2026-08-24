# PayLater Frontend

React + TypeScript + Vite SPA for Customer, Merchant, and Admin portals.

## Ports

| Surface | Host port | Notes |
|---------|-----------|--------|
| Jenkins | **8080** | Do not use for PayLater |
| PayLater UI (Docker) | **8081** | nginx → container `:80` |
| Vite dev | **5173** | `npm run dev` |
| API Gateway | **9090** | All browser API calls |

## Local development

```bash
cp .env.example .env   # VITE_API_BASE_URL=http://localhost:9090
npm install
npm run dev
```

Ensure the backend gateway is running on `:9090` (see `../paylater-backend/README.md`).

## Docker (with backend Compose)

From `paylater-backend/`:

```bash
docker compose up --build -d
```

UI: **http://localhost:8081**  
The image is built with `VITE_API_BASE_URL=/api`; nginx proxies `/api/` to the `gateway` service.

## Environment

| Variable | Purpose |
|----------|---------|
| `VITE_API_BASE_URL` | Axios base URL (baked at Vite build time) |

Suitable for Kubernetes: inject at image build or serve a runtime config later; never hardcode container IPs.

## Scripts

```bash
npm run lint
npm run build
npm run preview
```
