FROM ubuntu:22.04

WORKDIR /app

COPY configs ./configs/
COPY ./.bin/app ./.bin/

RUN mkdir logs

ENTRYPOINT ["./.bin/app"]
