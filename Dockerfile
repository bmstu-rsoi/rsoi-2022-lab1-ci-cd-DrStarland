# Application configuration
# Be sure to deploy your application with a config.production.json config file containing the correct values for your production environment and set the GOYAVE_ENV environment variable to production.

# Ensure that the app.environment entry is production.
# The server.host entry should be 0.0.0.0 if you want to open access to your service from anywhere. If you're using Apache or Nginx as a proxy on the same machine, keep it at 127.0.0.1 so the server will only be accessible through the proxy.
# Change the server.domain entry to your domain name, if you use one.
# The server.port and server.httpsPort will very likely require a change. Most of the time, you need 80 and 443 respectively.
# If you use https, be sure to provide the paths to your server.tls.cert and server.tls.key. Learn more here.
# server.debug must be set do false. You don't want anyone to get important information about your internal errors and therefore your code when an error occurs.
# Change your database connection credentials. database.autoMigrate should be set to false.
# #Build
# Of course, don't run your application with go run in production. Build your application using go build and deploy the executable, alongside the config files and resources directory.

# # The following Dockerfile is an example of a goyave application called docker-goyave:
# FROM golang:1.19

# # Set the Current Working Directory inside the container
# WORKDIR /home/osboxes/prog/rsoi-2022-lab1-ci-cd-DrStarland/probagoyave

# # Copy everything from the current directory to the PWD (Present Working Directory) inside the container
# COPY . .

# # Download all the dependencies
# RUN go get -d -v ./...

# # Install the package
# RUN go install -v ./...

# # This container exposes port 8080 to the outside world
# EXPOSE 8080

# # Run the executable
# CMD ["probagoyave"]


# syntax=docker/dockerfile:1

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.19-alpine

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY . ./

# compile application
RUN go build -o /godocker
 
# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8080

# command to be used to execute when the image is used to start a container
CMD [ "/godocker" ]
