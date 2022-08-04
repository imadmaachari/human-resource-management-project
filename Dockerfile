# Dockerfile References: https://docs.docker.com/engine/reference/builder/

## We'll choose the incredibly lightweight
## Go alpine image to work with
FROM golang:1.18 AS builder

# Add Maintainer Info
LABEL maintainer="Imad.M <imad.maachari@gmail.com>"

## We create an /app directory in which
## we'll put all of our project code
# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

## We want to build our application's binary executable
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

######## Start a new stage from scratch #######
## the lightweight scratch image we'll
## run our application within
FROM alpine:latest AS production
## We have to copy the output from our
## builder stage to our production stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080
## we can then kick off our newly compiled
## binary exectuable!!
CMD ["./main"]