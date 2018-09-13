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
RUN mkdir -p .bitcoin

RUN git checkout ${BITCOIND_VERSION}
RUN set -x && \
    ./autogen.sh && \
    ./configure --without-gui && \
    make -s -j4 && \
    make install

COPY ./containers/client/docker-entrypoint.sh ${BITCOIN_ROOT}/bitcoin/
RUN mkdir -p ${BITCOIN_ROOT}/.bitcoin
ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["bitcoin-cli", "help"]
