FROM ubuntu:16.04

RUN  \
  apt-get update && \
  apt-get -y install build-essential libtool autotools-dev automake pkg-config libssl-dev libevent-dev bsdmainutils python3 \
  libboost-system-dev libboost-filesystem-dev libboost-chrono-dev libboost-program-options-dev libboost-test-dev libboost-thread-dev && \
  apt-get -y install software-properties-common && \
  add-apt-repository ppa:bitcoin/bitcoin && \
  apt-get update && \
  apt-get -y install libdb4.8-dev libdb4.8++-dev && \
  apt-get -y install libqt5gui5 libqt5core5a libqt5dbus5 qttools5-dev qttools5-dev-tools libprotobuf-dev protobuf-compiler && \
  apt-get -y install git wget tar

RUN mkdir -p ~/app
WORKDIR /app
RUN git clone https://github.com/bitcoin/bitcoin.git

WORKDIR /app/bitcoin
RUN  git checkout "${BITCOIND_VERSION}"
RUN ./contrib/install_db4.sh `pwd`

RUN export BDB_PREFIX='/app/bitcoin/db4'
RUN \
  ./autogen.sh && \
  ./configure --without-gui BDB_LIBS="-L${BDB_PREFIX}/lib -ldb_cxx-4.8" BDB_CFLAGS="-I${BDB_PREFIX}/include" && \
  make -s -j4

