Simple Multi-Stage Dockerfile :
Using multi-stage Dockerfiles, we can pick apart the tasks of building and running our Go applications into different stages. Typically, we start off with a large image which includes all of the necessary dependencies, packages, etc. needed to compile the binary executable of our Go application. This would be classed as our builder stage.

We then take a far more lightweight image for our run stage which includes only what is absolutely needed in order to run a binary executable. This would typically be classed as a production stage or something similar.
## We use the larger image which includes
## all of the dependencies that we need to
## compile our program
FROM bigImageWithEverything AS Builder
RUN go build -o main ./...

## We then define a secondary stage which
## is built off a far smaller image which
## has the absolute bare minimum needed to
## run our binary executable application
FROM LightweightImage AS Production
CMD ["./main"]
By doing it this way, we benefit from a consistent build stage and we benefit from having absolutely tiny images in which our application will run in a production environment.

$ docker build -t golang-app .
$ docker run -it --rm --name running-app golang-app