FROM alpine:latest

ARG version=1.0.0

LABEL maintainer="Yohei Noguchi" \
      description="expanded wc (word count)"

RUN    adduser -D oilio \
    && apk --no-cache add --update --virtual .builddeps curl tar \
    #&& curl -s -L -O https://github.com/enoguch/oilio/releases/download/v${version}/oilio-${version}_linux_amd64.tar.gz \
    && curl -s -L -o oilio-${version}_linux_amd64.tar.gz https://www.dropbox.com/sh/qclkbrdf72m7os7/AACLckuZOa34GgeBT7gyin1Oa?dl=0 \
    && tar xfz oilio-${version}_linux_amd64.tar.gz     \
    && mv oilio-${version} /opt                        \
    && ln -s /opt/oilio-${version} /opt/oilio        \
    && ln -s /opt/oilio/oilio /usr/local/bin/oilio \
    && rm oilio-${version}_linux_amd64.tar.gz          \
    && apk del --purge .builddeps

ENV HOME="/home/oilio"

WORKDIR /home/oilio
USER    oilio

ENTRYPOINT [ "/opt/oilio/oilio" ]
