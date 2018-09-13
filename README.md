Please set the following variable to .evn
```
RPC_USER
RPC_AUTH
RPC_PASSWORD
BITCOIND_VERSION
RPC_BIND
RPC_ALLOWIP
```

Containers build!
```
dcc build
```

```
dcc up -d node
```

```
dcc up client
```

RPC access
```
dcc run --rm client bitcoin-cli <cmd>
```
