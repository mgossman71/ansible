FROM alpine
# ENV PATH=$PATH:/root
RUN mkdir /code
WORKDIR /code
RUN apk add go && \
    apk add git && \
    apk add ansible && \
    apk add go
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux
# COPY main.go .
# RUN go build main.go
# CMD mkdir .kube
# COPY config .kube/config
# ENTRYPOINT [ "main" ] 
# EXPOSE 8080
