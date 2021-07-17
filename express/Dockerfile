FROM node:16-alpine3.11

RUN apk add --no-cache inotify-tools

COPY ./package.json /app/package.json
COPY ./yarn.lock /app/yarn.lock
COPY ./tsconfig.json /app/tsconfig.json
COPY ./nodemon.json /app/nodemon.json

VOLUME /app/src

WORKDIR /app

RUN ["yarn"]

EXPOSE 3000

ENTRYPOINT ["yarn", "start"]

