# mongof

Simple command line utility to tail the mongo oplog from your terminal.

## Install

Run:

``` terminal
$ go install github.com/pjvds/mongof
```

## Examples

### get all docs

``` terminal
$ mongof
{
    "ns": "",
    "o": {},
    "op": "n",
    "ts": 6135382171273134081
}
{
    "ns": "opchat.threads",
    "o": {
        "_id": "55254194b479542277d3eaf4",
        "by_user_id": "1",
        "create_at": "2015-04-08T16:56:20.658+02:00",
        "message": "foobar",
        "thread_id": "1",
        "to_user_id": "2"
    },
    "op": "i",
    "ts": 6135382171273134082
}
{
    "ns": "opchat.threads",
    "o": {
        "_id": "55254196b479542277d3eaf5",
        "by_user_id": "1",
        "create_at": "2015-04-08T16:56:22.705+02:00",
        "message": "foobar",
        "thread_id": "1",
        "to_user_id": "2"
    },
    "op": "i",
    "ts": 6135382179863068673
}
```

### get everything execpt time no-ops

``` terminal
$ mongof -query='{"op":{"$ne":"n"}}'
{
    "ns": "opchat.threads",
    "o": {
        "_id": "55254194b479542277d3eaf4",
        "by_user_id": "1",
        "create_at": "2015-04-08T16:56:20.658+02:00",
        "message": "foobar",
        "thread_id": "1",
        "to_user_id": "2"
    },
    "op": "i",
    "ts": 6135382171273134082
}
{
    "ns": "opchat.threads",
    "o": {
        "_id": "55254196b479542277d3eaf5",
        "by_user_id": "1",
        "create_at": "2015-04-08T16:56:22.705+02:00",
        "message": "foobar",
        "thread_id": "1",
        "to_user_id": "2"
    },
    "op": "i",
    "ts": 6135382179863068673
}
```

```