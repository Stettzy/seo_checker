# SEOBuddy Production Deployment

This branch contains the production configuration for deploying SEOBuddy to your server.

## Prerequisites

1. Build the frontend and backend locally:

```bash
# Build frontend
cd frontend
npm install
npm run build

# Build backend
cd ../backend
go build -o main
```

2. Make sure you have the following files ready:
   - `frontend/dist/` directory (from frontend build)
   - `backend/main` binary (from backend build)

## Deployment

1. Clone this branch on your server:
```bash
git clone -b production https://github.com/stettzy/seo-checker.git
cd seo-checker
```

2. Update your domain in:
   - `docker-compose.prod.yml` (VITE_API_URL)
   - `nginx/conf.d/app.conf` (server_name)
   - `init-letsencrypt.sh` (domains array)

3. Initialize SSL certificates:
```bash
chmod +x init-letsencrypt.sh
./init-letsencrypt.sh
```

4. Start the services:
```bash
docker compose -f docker-compose.prod.yml up -d
```

## Notes

- The production setup assumes you've already built the frontend and backend locally
- SSL certificates will auto-renew every 12 hours
- Nginx configuration includes security headers and optimizations
- Both frontend and backend are served behind Nginx reverse proxy 