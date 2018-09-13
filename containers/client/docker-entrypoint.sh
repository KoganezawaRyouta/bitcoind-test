#!/bin/bash
set -e

if [[ ! -s "${BITCOIN_ROOT}/.bitcoin/bitcoin.conf" ]]; then
    echo "Creating bitcoin.conf"
    cat <<-EOF > "${BITCOIN_ROOT}/.bitcoin/bitcoin.conf"
rpcconnect=${RPC_BIND}
rpcport=18332
rpcpassword=${RPC_PASSWORD}
rpcuser=${RPC_USER}
rpcauth=${RPC_AUTH}
datadir=${BITCOIN_ROOT}
EOF
    echo "Initialization completed successfully"
fi

exec "$@"
