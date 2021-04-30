# lightweight container for go
FROM golang:alpine

# update container's packages and install git
RUN apk update && apk add --no-cache git

# set env
ENV MYSQL_CONN_STRING=""
ENV REDIS_CONN_STRING=""
ENV REDIS_CONN_PASS_STRING=""

ENV PATH_STATIC_ENG     = "./public/images/engineers"
ENV PATH_UPLOAD_ENG     = "public/images/engineers/"
ENV PUBLIC_SHARE_IMG    = "/public/engineers"
ENV PUBLIC_UPLOAD_DB    = "public/engineers/"

ENV ACCESS_SECRET = jdnfksdmfksd
ENV REFRESH_SECRET = mcmvmkmsdnfsdmfdsjf

# set /app to be the active directory
WORKDIR /app

# copy all files from outside container, into the container
COPY . .

# download dependencies
RUN go mod tidy

# build binary
RUN go build -o binary

# set the entry point of the binary
ENTRYPOINT ["/app/binary"]
