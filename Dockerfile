# Sử dụng Go phiên bản chính thức
FROM golang:1.20

# Tạo thư mục làm việc bên trong container
WORKDIR /app

# Copy các file module
COPY go.mod go.sum ./
RUN go mod download

# Copy mã nguồn
COPY . .

# Build ứng dụng
RUN go build -o main .

# Expose cổng ứng dụng
EXPOSE 8080

# Chạy ứng dụng
CMD ["./main"]
