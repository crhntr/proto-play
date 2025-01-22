# Playground for Postgres+Protobuf

```sh
goose -dir migrations postgres "host=${PGHOST} dbname=${PGDATABASE} ${PGQUERYSTRING} user=${PGUSER}" up
```

GRP CURL
```sh
grpcurl -d '{"query": {"un": {"value": "EXAMPLE"}}}' -plaintext -proto ./proto/api/playground/v1/playground.proto localhost:8767 api.playground.v1.StoreService/Exists
```
