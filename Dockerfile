FROM golang:1.16.3 AS builder
RUN mkdir /clusterGetnodeip
WORKDIR /clusterGetnodeip
COPY ./clusterGetnodeip/* /clusterGetnodeip/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


FROM haproxy:alpine
# Get confd
ENV CONFD_VERSION 0.16.0
ADD https://github.com/kelseyhightower/confd/releases/download/v${CONFD_VERSION}/confd-${CONFD_VERSION}-linux-amd64 /usr/bin/confd

# Get kubectl
RUN apk add --no-cache curl bash bash-completion vim
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl && mv ./kubectl /usr/local/bin/kubectl && chmod +x /usr/local/bin/kubectl

COPY confd/ /etc/confd/
COPY endpoint.sh /endpoint.sh
COPY haproxy.cfg /usr/local/etc/haproxy/haproxy.cfg

# 此处添加小工具
COPY --from=builder /clusterGetnodeip/clusterGetnodeip /usr/bin/clusterGetnodeip

RUN chmod +x /usr/bin/confd /endpoint.sh /usr/bin/clusterGetnodeip
#RUN chmod +x  /endpoint.sh /usr/bin/clusterGetnodeip

CMD /endpoint.sh
