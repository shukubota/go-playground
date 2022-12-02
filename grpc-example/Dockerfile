FROM golang:1.19

RUN apt-get update
RUN apt-get install sudo
RUN apt-get install -y protobuf-compiler
RUN curl -sL https://deb.nodesource.com/setup_16.x | sudo -E bash - && sudo apt install -y nodejs
RUN npm -g install protoc-gen-grpc-web
RUN npm -g install protoc-gen-ts
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
COPY ./run.sh /var/work/
COPY ./protobuf /var/work/protobuf
WORKDIR /var/work/
CMD ["bash", "run.sh"]