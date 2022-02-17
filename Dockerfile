FROM golang:1.17-alpine AS build

WORKDIR /src/
COPY . .
RUN CGO_ENABLED=0 go install .

RUN useradd -u 1234 nonroot

FROM scratch
COPY --from=build /go/bin/vecurs /bin/vecurs
USER nonroot
ENTRYPOINT ["/bin/vecurs"]
