# Playground for Postgres+Protobuf

```sh
goose -dir migrations postgres "host=${PGHOST} dbname=${PGDATABASE} ${PGQUERYSTRING} user=${PGUSER}" up
```

GRP CURL
```sh
grpcurl -d '{"query": {"un": {"value": "EXAMPLE"}}}' -plaintext -proto ./proto/api/playground/v1/playground.proto localhost:8767 api.playground.v1.StoreService/Exists
```


Add these to get an unmarshal error.
```
insert into messages (id, content)
values  (1, '{"un": {"value": "EXAMPLE"}}'),
        (2, '{"un": {"value": "a"}}'),
        (3, '{"un": {"value": "b"}}'),
        (4, '{"pn": {"last": "l", "first": "f"}}'),
        (5, '{"pn": {"last": "ll", "first": "ff"}}'),
        (30, '{"pn": {"last": "ll", "first": "f2"}}'),
        (34, '{"pn": {"last": "ll", "first": "f3"}}'),
        (38, '{"pn": {"last": "ll", "first": "f3", "middle": ["m1", "m2"]}}'),
        (41, '{"pn": {"last": "ll", "first": "f3", "middle": ["m1", "m2", "m3"]}}');
```
