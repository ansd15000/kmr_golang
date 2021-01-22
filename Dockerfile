FROM golang:1.11-alpine AS build

MAINTAINER sanghong,ahn <ansd1500@gmail.com>
LABEL "purpose"="Go Application Service Deployment"

WORKDIR /src/
COPY main.go go.* /src
RUN CGO_ENABLE=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]
