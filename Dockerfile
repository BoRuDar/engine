FROM node:slim
MAINTAINER Vitaly Kovalyshyn "v.kovalyshyn@webitel.com"

ENV VERSION 4
ENV NODE_TLS_REJECT_UNAUTHORIZED 0

COPY src /engine

WORKDIR /engine
RUN npm install && npm cache clear

EXPOSE 10022
ENTRYPOINT ["node", "server.js"]
