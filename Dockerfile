FROM alpine
ENV PATH=$PATH:/code
RUN mkdir /code
RUN mkdir ~/.ssh && touch ~/.ssh/id_rsa && chmod -R 600 ~/.ssh 
WORKDIR /code
RUN apk add go && \
    apk add git && \
    apk add ansible && \
    apk add openssh && \
    apk add go
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux
RUN git clone https://github.com/mgossman71/playbook-centos-base.git
WORKDIR /code/playbook-centos-base
COPY ansible.go .
RUN go build ansible.go
# ENTRYPOINT [ "./ansible" ] 
EXPOSE 8080
