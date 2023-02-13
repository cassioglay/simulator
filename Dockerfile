FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

#PLUGIN PARA TRABALHAR COM O KAFKA
RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y 

CMD ["tail", "-f", "/dev/null"]