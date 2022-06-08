FROM node:14.18-alpine as web

WORKDIR /opt/workflow

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk update && \
    apk add --no-cache git && \
    rm -rf /var/cache/apk/* /tmp/* /var/tmp/* $HOME/.cache
RUN git clone https://gitee.com/yllan/daisy_web.git

WORKDIR daisy_web

RUN npm install -g cnpm --registry=https://registry.npm.taobao.org
RUN cnpm install
RUN echo $'# just a flag\n\
ENV = \'production\'\n\n\
# base api\n\
VUE_APP_BASE_API = \'\''\
> .env.production
RUN npm run build:prod

FROM golang:1.15 AS build

WORKDIR /opt/workflow/daisy
COPY . .
ARG GOPROXY="https://goproxy.cn"
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o daisy .

FROM alpine AS prod

MAINTAINER xmlKevin

WORKDIR /opt/workflow/daisy

COPY --from=build /opt/workflow/daisy/daisy /opt/workflow/daisy/
COPY config/ /opt/workflow/daisy/default_config/
COPY template/ /opt/workflow/daisy/template/
COPY docker/entrypoint.sh /opt/workflow/daisy/
RUN mkdir -p logs static/uploadfile static/scripts static/template

RUN chmod 755 /opt/workflow/daisy/entrypoint.sh
RUN chmod 755 /opt/workflow/daisy/daisy

COPY --from=web /opt/workflow/daisy_web/web /opt/workflow/daisy/static/web
COPY --from=web /opt/workflow/daisy_web/web/index.html /opt/workflow/daisy/template/web/

RUN mv /opt/workflow/daisy/static/web/static/web/* /opt/workflow/daisy/static/web/
RUN rm -rf /opt/workflow/daisy/static/web/static

EXPOSE 8002
VOLUME [ "/opt/workflow/daisy/config" ]
ENTRYPOINT [ "/opt/workflow/daisy/entrypoint.sh" ]
