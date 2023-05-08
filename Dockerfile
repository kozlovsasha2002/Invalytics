FROM golang:1.19

RUN mkdir D:\\Aleksandr\\Golang\\projects\\Invalytics

ADD . /d/Aleksandr/Golang/projects/Invalytics

WORKDIR /mnt/d/Aleksandr/Golang/projects/Invalytics

RUN go build -o main /mnt/d/Aleksandr/Golang/projects/Invalytics/app/cmd

CMD ["/mnt/d/Aleksandr/Golang/projects/Invalytics/app/cmd/main"]