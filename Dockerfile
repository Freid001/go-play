FROM ubuntu:14.04

LABEL maintainer=freid

RUN apt-get update && apt-get install --no-install-recommends -y \
    ca-certificates \
    curl \
    mercurial \
    git-core
RUN curl -s https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz | tar -v -C /usr/local -xz

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH /usr/local/go/bin:/go/bin:/usr/local/bin:$PATH

COPY . /go/src/github.com/freid001/go-play

WORKDIR /go/src/github.com/freid001/go-play

CMD ["bash"]