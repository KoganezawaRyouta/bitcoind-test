FROM ubuntu:16.04

RUN  \
  apt-get update && \
  apt-get -y install build-essential libtool autotools-dev automake pkg-config libssl-dev libevent-dev bsdmainutils \
  python3 \
  libboost-all-dev \
  libboost-system-dev libboost-filesystem-dev libboost-chrono-dev libboost-program-options-dev libboost-test-dev libboost-thread-dev \
  software-properties-common && \
  add-apt-repository ppa:bitcoin/bitcoin && \
  apt-get update && \
  apt-get -y install libdb4.8-dev libdb4.8++-dev git \
  libqt5gui5 libqt5core5a libqt5dbus5 qttools5-dev qttools5-dev-tools libprotobuf-dev protobuf-compiler libqrencode-dev

RUN git clone https://github.com/bitcoin/bitcoin.git
WORKDIR /bitcoin
RUN  git checkout $BITCOIND_VERSION
