FROM golang:1.18-alpine AS development

ENV PROJECT_PATH=/payment-service
ENV PATH=$PATH:$PROJECT_PATH/build
ENV CGO_ENABLED=0
ENV GO_EXTRA_BUILD_ARGS="-a -installsuffix cgo"

RUN apk add --no-cache ca-certificates tzdata make git bash

RUN mkdir -p $PROJECT_PATH
COPY . $PROJECT_PATH
WORKDIR $PROJECT_PATH

RUN go mod download
RUN make

FROM alpine:3.11.2 AS production

RUN apk --no-cache add ca-certificates tzdata
COPY --from=development /payment-service/build/payment-service /usr/bin/payment-service
ENTRYPOINT ["/usr/bin/payment-service"]