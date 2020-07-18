FROM gradle:latest

RUN mkdir /app
WORKDIR /app
COPY ./app/ /app/
ENTRYPOINT ["gradle", "run"]

EXPOSE 8081 8082
