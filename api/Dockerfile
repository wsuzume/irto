FROM golang:1.14.4

ENV ENVIRONMENT=debug

WORKDIR /go/src/work
COPY src .
RUN go build -o ${ENVIRONMENT}

WORKDIR /app
# make directories, copy binary to /app/bin, and remove source code
RUN mkdir views \
    && mkdir assets \
    && mkdir bin \
    && cp /go/src/work/${ENVIRONMENT} bin \
    && rm -rf /go/src/work
# copy files to /app/views
COPY views views
#
COPY assets assets

