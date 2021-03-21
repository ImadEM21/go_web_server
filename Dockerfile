FROM golang:1.16 as Builder
WORKDIR /myapi
COPY . ./
RUN useradd api
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o api
FROM scratch
COPY --from=Builder /myapi/api /usr/bin/api
COPY --from=Builder /etc/passwd /etc/
USER api
EXPOSE 3000
ENTRYPOINT ["/usr/bin/api"]