# Golang 베이스 이미지
FROM golang:1.23.1 AS builder

# 작업 디렉토리 설정
WORKDIR /app

# Go 모듈 및 애플리케이션 복사
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 애플리케이션 빌드
RUN go build -o main .

# 실행 단계
FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/uploads ./uploads

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && apt-get clean

# 포트 노출
EXPOSE 8080

# 애플리케이션 실행
CMD ["./main"]
