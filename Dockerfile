FROM golang:1.15.7

EXPOSE 8081
WORKDIR /app

COPY . /app
 
RUN go get -d github.com/gorilla/mux

# RUN go install 
RUN go build ./app.go
 
CMD ["./app"]