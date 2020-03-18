#-------------------------building docker from golang image--------------------------#

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Anisur Rahman <emailforanis@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

# Build the Go app
RUN go mod tidy
RUN go mod vendor
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

#-------------------------building docker image just from the binary--------------------------#
#FROM ubuntu
#COPY first_api_project /bin
#CMD ["first_api_project"]