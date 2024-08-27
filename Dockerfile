FROM alpine:edge as builder
LABEL stage=go-builder
WORKDIR /app/
# RUN apk add --no-cache bash curl gcc git go musl-dev
RUN sed -i 's|https://dl-cdn.alpinelinux.org/alpine/|https://mirrors.aliyun.com/alpine/|g' /etc/apk/repositories && \
    apk add --no-cache bash curl gcc git go musl-dev openssl
# 设置 Go 模块代理
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN bash build.sh release docker

FROM alpine:edge

ARG INSTALL_FFMPEG=false
LABEL MAINTAINER="i@nn.ci"

WORKDIR /opt/alist/

# RUN apk update && \
#     apk upgrade --no-cache && \
#     apk add --no-cache bash ca-certificates su-exec tzdata; \
#     [ "$INSTALL_FFMPEG" = "true" ] && apk add --no-cache ffmpeg; \
#     rm -rf /var/cache/apk/*
# 换源
RUN sed -i 's|https://dl-cdn.alpinelinux.org/alpine/|https://mirrors.aliyun.com/alpine/|g' /etc/apk/repositories && \
    apk update && \
    apk upgrade --no-cache && \
    apk add --no-cache bash ca-certificates su-exec tzdata; \
    [ "$INSTALL_FFMPEG" = "true" ] && apk add --no-cache ffmpeg; \
    rm -rf /var/cache/apk/*


COPY --from=builder /app/bin/alist ./
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh && /entrypoint.sh version

ENV PUID=0 PGID=0 UMASK=022
VOLUME /opt/alist/data/
EXPOSE 5244 5245
CMD [ "/entrypoint.sh" ]