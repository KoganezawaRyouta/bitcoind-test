#!/bin/bash
set -e

if [[ ! -s "${BITCOIN_ROOT}/.bitcoin/bitcoin.conf" ]]; then
    echo "Creating bitcoin.conf"
    cat <<-EOF > "${BITCOIN_ROOT}/.bitcoin/bitcoin.conf"
server=1
testnet=3
txindex=1
printtoconsole=1
rpcpassword=${RPC_PASSWORD}
rpcuser=${RPC_USER}
rpcauth=${RPC_AUTH}
rpcport=18332
rpcbind=${RPC_BIND}
rpcallowip=${RPC_ALLOWIP}
datadir=${BITCOIN_ROOT}/.bitcoin
listen=1
EOF
    echo "Initialization completed successfully"
fi

exec "$@"
