FROM golang:1.17-alpine AS build

WORKDIR /src/
COPY . .
RUN go mod download && CGO_ENABLED=0 go install .

ARG USERNAME=nonroot
ARG USER_UID=1000
ARG USER_GID=${USER_UID}
RUN addgroup -g ${USER_GID} -S ${USERNAME} \
    && adduser -D -u ${USER_UID} -S ${USERNAME} -s /bin/sh ${USERNAME}

FROM scratch

COPY --from=build /go/bin/bcvcurs /bin/bcvcurs
USER ${USERNAME}
ENTRYPOINT ["/bin/bcvcurs"]
