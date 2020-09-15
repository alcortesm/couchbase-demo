A demo of the couchbase Go client.

Run a local Couchbase cluster with:

```
; docker run -t --name db -p 8091-8094:8091-8094 -p 11210:11210 -v /tmp/couchbase_demo:/opt/couchbase/var couchbase/server-sandbox:6.5.0
```

Then try to connect and run some queries:

```
; go run .
gocb version: v2.1.5
cluster ready
ping result: {
  "version": 2,
  "sdk": "gocb/v2.1.5 gocbcore/v9.0.5",
  "id": "ea808343-66ad-4b29-ac8f-475100f21ff7",
  "services": {
    "mgmt": [
      {
        "id": "94e7ccc7-8f3d-467a-bc37-e9a3067a51ca",
        "remote": "http://localhost:8091",
        "state": "ok",
        "latency_us": 7601111
      }
    ],
    "query": [
      {
        "id": "dbca3ddb-920b-4cb7-83e4-23f522ccd974",
        "remote": "http://localhost:8093",
        "state": "ok",
        "latency_us": 2516301
      }
    ],
    "search": [
      {
        "id": "0d28c9aa-cc28-4663-9650-3f3cdc72d945",
        "remote": "http://localhost:8094",
        "state": "ok",
        "latency_us": 1752641
      }
    ]
  }
}
cluster diagnostics: {
  "version": 2,
  "sdk": "gocb/v2.1.5",
  "id": "408026f7-97f8-41cf-aef1-b9e58fec1721",
  "services": {
    "kv": [
      {
        "id": "0xc000064820",
        "last_activity_us": 511200382,
        "remote": "localhost:11210",
        "local": "[::1]:51424",
        "state": "connected"
      }
    ]
  },
  "state": "online"
}
2020/09/15 19:54:28 getting result metadata: the result must be closed before accessing the meta-data
exit status 1
```
