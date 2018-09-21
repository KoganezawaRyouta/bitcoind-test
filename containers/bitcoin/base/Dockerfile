FROM ubuntu:16.04

RUN set -x && \
  rm -rf /var/lib/apt/lists/* && \
  apt-get clean && \
  apt-get update

RUN set -x && \
  apt-get -y install build-essential libtool autotools-dev automake pkg-config libssl-dev libevent-dev bsdmainutils python3 \
  libboost-system-dev libboost-filesystem-dev libboost-chrono-dev libboost-program-options-dev libboost-test-dev libboost-thread-dev && \
  apt-get -y install software-properties-common && \
  add-apt-repository ppa:bitcoin/bitcoin && \
  apt-get update && \
  apt-get -y install libdb4.8-dev libdb4.8++-dev && \
  apt-get -y install libqt5gui5 libqt5core5a libqt5dbus5 qttools5-dev qttools5-dev-tools libprotobuf-dev protobuf-compiler && \
  apt-get -y install git wget tar

ONBUILD RUN echo "See you later"
