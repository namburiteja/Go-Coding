# Multi-stage build: Vite SPA + nginx (SPA routing + /api → gateway).
# Build from paylater-frontend/:
#   docker build -t paylater-frontend:local \
#     --build-arg VITE_API_BASE_URL=/api .

FROM node:24-alpine AS build

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .

# Browser calls same-origin /api; nginx proxies to the gateway service.
ARG VITE_API_BASE_URL=/api
ENV VITE_API_BASE_URL=$VITE_API_BASE_URL

RUN npm run build

FROM nginx:1.27-alpine

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
