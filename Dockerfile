FROM golang:alpine AS build
RUN apk add --update --no-cache ca-certificates git
ADD . /src
RUN cd /src && CGO_ENABLED=0 go build -o cheshire-east-bin-collection-ics

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/cheshire-east-bin-collection-ics /
EXPOSE 8080
ENV UPRN "100012357047"
ENTRYPOINT ["/cheshire-east-bin-collection-ics"]
