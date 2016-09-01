FROM marcbachmann/libvips

RUN apt-get update && apt-get install --no-install-recommends -y \
    ca-certificates \
    curl \
    mercurial \
    git-core
RUN curl -s https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz | tar -v -C /usr/local -xz

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH /usr/local/go/bin:/go/bin:/usr/local/bin:$PATH

ADD . /go/src/github.com/carloct/imgrazor

RUN go install github.com/carloct/imgrazor

ENTRYPOINT /go/bin/imgrazor

EXPOSE 80