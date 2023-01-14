FROM ubuntu:22.04

WORKDIR /app

COPY configs ./configs/
COPY ./.bin/main ./.bin/

CMD ["./.bin/main"]
