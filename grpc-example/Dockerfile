FROM golang:1.22

RUN apt-get update
RUN apt-get install sudo
RUN apt-get install -y protobuf-compiler
RUN apt install -y nodejs npm
RUN npm install n -g && n stable
RUN apt purge -y nodejs npm && apt autoremove -y

RUN npm -g install protoc-gen-grpc-web
RUN npm -g install protoc-gen-ts
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
COPY ./run.sh /var/work/
COPY ./protobuf /var/work/protobuf
WORKDIR /var/work/
CMD ["bash", "run.sh"]