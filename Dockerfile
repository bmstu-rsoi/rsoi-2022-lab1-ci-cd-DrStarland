# FROM golang:1.16-buster AS build

# WORKDIR /app

# RUN mkdir probagoyave
# COPY ./probagoyave ./probagoyave
# WORKDIR /app/probagoyave
# RUN go mod download

# RUN go build -o /docker-gs-ping

# ## Deploy
# FROM gcr.io/distroless/base-debian10

# WORKDIR /

# COPY --from=build /docker-gs-ping /docker-gs-ping

# EXPOSE 8080

# USER nonroot:nonroot

# ENTRYPOINT ["/docker-gs-ping"]

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.19-alpine

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY ./probagoyave/go.mod ./probagoyave/go.sum ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY ./probagoyave ./

# compile application
RUN go build -o /godocker
 
# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8080

# command to be used to execute when the image is used to start a container
CMD [ "/godocker" ]