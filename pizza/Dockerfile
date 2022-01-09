FROM golang:1.17.5-bullseye
WORKDIR /app
RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
CMD ["/bin/air"]
