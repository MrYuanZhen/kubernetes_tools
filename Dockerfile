FROM golang:1.16.3
RUN mkdir /clusterGetnodeip
WORKDIR /clusterGetnodeip
COPY ./* /clusterGetnodeip/
RUN go install
