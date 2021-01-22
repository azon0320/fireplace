# Do not build directly. Use `make amd64-docker` instead (at this file's directory)
FROM alpine:3.12

RUN mkdir -p /var/app/fireplace

# copy bin dist to /tmp
COPY bin/. /var/app/fireplace

# fireplace server uses ENVs below
ENV MONGO_URL=root:rootpw@mongodb:27017
ENV SERVER_ADDRESS=:8999

# fireplace server serves default
EXPOSE 8999

WORKDIR /var/app/fireplace

ENTRYPOINT /var/app/fireplace/fireplace-server

