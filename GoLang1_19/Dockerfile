FROM golang:1.19.3-alpine

WORKDIR /home

RUN apk add --no-cache git gcc musl-dev make && \
	git clone https://github.com/altoliaw/SecuEyesDecoder.git && \
	cd SecuEyesDecoder && \
	make archiveDLL