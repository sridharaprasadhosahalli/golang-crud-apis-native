
FROM golang:1.23.4 AS build-stage

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-golang-crud-apis

# Run the tests in the container
FROM build-stage AS run-test-stage

RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-golang-crud-apis /docker-golang-crud-apis

#expose 8080 to outside
EXPOSE 8080

#to be run as non root
USER nonroot:nonroot

ENTRYPOINT ["/docker-golang-crud-apis"]