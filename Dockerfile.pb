FROM golang:1.19-bullseye

RUN apt update -y && apt upgrade -y
RUN apt-get update -y && apt-get upgrade -y
RUN apt install make

COPY . /fs3
COPY keys/id_ecdsa.pub /keys/id_ecdsa.pub
WORKDIR /fs3

RUN go mod download
RUN make server -B

CMD [ "server/app/server" ]