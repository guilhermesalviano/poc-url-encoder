FROM golang:1.19.4-alpine3.17 as builder

RUN apk update && apk add --no-cache make git openssh-client ca-certificates bash

ENV DEPNOLOCK=1
ENV RELEASE_TAG 'base'
ENV TZ 'America/Sao_Paulo'
ENV SERVICE_PORT 3333

WORKDIR /go/src/url-encoder

COPY . .

RUN ssh-keygen -b 2048 -t rsa -f /root/.ssh/id_rsa -q -N ""
RUN echo -e ${GITHUB_PRIVATE_KEY} > /root/.ssh/id_rsa && rm -rf /root/.ssh/id_rsa.pub
RUN ssh -tt -o StrictHostKeyChecking=no -i /root/.ssh/id_rsa git@github.com || true

RUN git config --global --add url.'git@github.com:'.insteadOf 'https://github.com/'

RUN make deps

RUN make build

RUN mv url-encoder /usr/bin/

CMD [ "/usr/bin/url-encoder" ]

EXPOSE 3333

#STOPSIGNAL SIGINT

#HEALTHCHECK --interval=30s \
#    CMD wget --spider http://localhost:3333/health/ || exit 1
