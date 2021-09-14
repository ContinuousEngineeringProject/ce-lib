FROM continuousengineeringproject/ce-base-imgae:0.0.1

WORKDIR /

COPY ./

RUN go mod download

ENTRYPOINT go run