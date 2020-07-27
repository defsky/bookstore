FROM debian:latest

RUN mkdir /app
WORKDIR /app

ADD user-service /app/user-service
ADD conf ./conf/

CMD [ "/app/user-service" ]
