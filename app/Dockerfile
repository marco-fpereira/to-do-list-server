FROM golang:1.22.4-alpine as builder
WORKDIR /app
ENV CGO_ENABLED=0
ENV GO111MODULE="on"
ENV GOOS=linux
COPY ./ ./
RUN go mod download
RUN go build -o ./app ./app.go
RUN chmod +x ./app

FROM alpine:3.14 as application
WORKDIR /app
COPY --from=builder /app /app
RUN addgroup --gid 1001 -S gouser && adduser --uid 1001 -S gouser -G gouser
RUN chown -R gouser:gouser /app
USER gouser:gouser
EXPOSE 8080 8888 3000 50051
CMD ["./app"]
# cd ./app
# docker build -t "to-do-list-server:v1" --target application .
# cd ..
# kubectl apply -f ./infra/kubernetes.yml
# kubectl get pods
# kubectl describe pod to-do-list-server-{pod-hash}
