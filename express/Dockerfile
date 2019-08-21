FROM alpine:3.10

RUN apk add --no-cache inotify-tools
RUN apk add --no-cache nodejs-current npm

COPY ./package.json /app/package.json
COPY ./package-lock.json /app/package-lock.json
COPY ./tsconfig.json /app/tsconfig.json
COPY ./nodemon.json /app/nodemon.json

VOLUME /app/src

WORKDIR /app

RUN ["npm", "i"]

EXPOSE 3000

ENTRYPOINT ["npm", "start"]

