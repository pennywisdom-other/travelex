FROM golang:1.8 as build
WORKDIR /go/src/app
COPY . .
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."
RUN go test -v
RUN go build
RUN godoc app
RUN ls

FROM golang:1.8
COPY --from=build /go/src/app/app /go/src/app/app
RUN chmod +x /go/src/app/app
EXPOSE 8080
WORKDIR /go/src/app
CMD ["./app"]
