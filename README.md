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

Execute bitcoind server 
```
dcc up -d node
```

Execute rpc client
```
dcc run --rm client bitcoin-cli <cmd>
```
