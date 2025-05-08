package e2e

var Cfg = `
postgres:
  addr: postgres:5432
  user_name: balance
  password: balance
  database: balance
  schema: public
  max_idle_connections: 10
  max_open_connections: 100
  connection_max_lifetime: 1h
  return_utc: true
  migrations_dir: ./internal/balance/migrations/postgres

use_cases:
  report_revenue_export:
    folder: /tmp/report_revenue_export
    ttl: 1h
    url: /data/report_revenue_export/
    zip: true
    
jwt:
  secret: "secret_key"
  ttl: 24h

http:
  addr: 0.0.0.0:8080
  api_prefix: /api
  handler_timeout: 10s
  read_timeout: 10s
  write_timeout: 10s
  swagger:
    enabled: true
    path: /swagger
  cors:
    allowed_origins:
      - "*"
  tls:
    enabled: false
    cert_file: ./certs/cert.pem
    key_file: ./certs/key.pem
  

prometheus:
  port: 8081
  path: /metrics

grpc:
  addr: 0.0.0.0:8083
  keepalive:
    timeout: 10s
    max_connection_idle: 10s
    max_connection_age: 10s
    max_connection_age_grace: 10s

mode: dev

lagging:
  enabled: false

jaeger:
  host: localhost:6831
  service_name: balance-service
  log_spans: true
  
`
