# BUILD

FROM docker.io/library/golang:1.13 as build
WORKDIR /work
COPY . /work/
RUN CGO_ENABLED=0 go build -mod=vendor -o /tmp/app .

# RUN

FROM gcr.io/distroless/static
WORKDIR /work

COPY --from=build --chown=nonroot:nonroot /tmp/app /work/app

USER nonroot:nonroot
ENTRYPOINT ["/work/app"]
