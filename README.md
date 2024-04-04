# postgres-backup-service

A lightweight CRON job service designed to save PostgreSQL dumps locally.

## How to define where to store backups

Update the `volumes` property in the `compose.yaml` file.

```yaml
volumes:
    - <full-path-here>:/app/backups
```

## How to start service

```bash
docker compose up -d
```
