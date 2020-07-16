FROM ubuntu:18.04

RUN apt-get update \
    && apt-get install -y curl nodejs npm

# Install latest npm
RUN npm install -g n
RUN n stable
RUN apt-get purge -y nodejs npm

# Install parcel
RUN npm install -g parcel-bundler

# Install Elm
# https://github.com/elm/compiler/blob/master/installers/linux/README.md
WORKDIR /elm
RUN curl -L -o elm.gz https://github.com/elm/compiler/releases/download/0.19.1/binary-for-linux-64-bit.gz
RUN gunzip elm.gz \
    && chmod +x elm \
    && mv elm /usr/local/bin/

WORKDIR /
RUN rm -r /elm
