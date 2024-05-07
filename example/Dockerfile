# Sử dụng image golang:latest làm cơ sở
FROM golang:1.21.3

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Sao chép mã nguồn vào thư mục làm việc trong container
COPY . .

# Biên dịch mã nguồn Go thành một tập tin nhị phân
RUN go build -o main .

# Mở cổng 8080 để lắng nghe kết nối HTTP
EXPOSE 8080

# Chạy ứng dụng khi container được khởi chạy
CMD ["./main"]
