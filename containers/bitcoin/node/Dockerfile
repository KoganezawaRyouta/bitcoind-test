FROM rk2826/bitcoind-base

ARG BITCOIND_VERSION
ARG RPC_USER
ARG RPC_AUTH
ARG RPC_PASSWORD
ARG RPC_BIND
ARG RPC_ALLOWIP

ENV BITCOIN_ROOT=/root
ENV RPC_USER=${RPC_USER}
ENV RPC_PASSWORD=${RPC_PASSWORD}
ENV RPC_AUTH=${RPC_AUTH}
ENV RPC_BIND=${RPC_BIND}
ENV RPC_ALLOWIP=${RPC_ALLOWIP}

RUN mkdir -p ${BITCOIN_ROOT}
WORKDIR ${BITCOIN_ROOT}
RUN git clone https://github.com/bitcoin/bitcoin.git

WORKDIR ${BITCOIN_ROOT}/bitcoin
RUN git checkout ${BITCOIND_VERSION}
RUN set -x && \
    ./contrib/install_db4.sh `pwd` && \
    ./autogen.sh && \
    ./configure --without-gui BDB_LIBS="-L${BITCOIN_ROOT}/bitcoin/db4/lib -ldb_cxx-4.8" BDB_CFLAGS="-I${BITCOIN_ROOT}/bitcoin/db4/include" && \
    make -s -j4 && \
    make install

COPY ./containers/bitcoin/node/docker-entrypoint.sh ${BITCOIN_ROOT}/bitcoin/
RUN mkdir -p ${BITCOIN_ROOT}/.bitcoin
ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["bitcoind", "-reindex-chainstate"]
