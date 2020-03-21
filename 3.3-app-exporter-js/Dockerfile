FROM node:latest

RUN mkdir /app
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH
ADD ./app/package.json ./app/yarn.lock /app/
COPY ./app/ /app/

RUN yarn install
ENTRYPOINT ["node", "src/main.js"]

EXPOSE 8081
