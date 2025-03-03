# Huớng dẫn cài đặt

1. Cài đặt Docker và Docker Compose

2. Tạo file môi trường:

    ```bash
    cp .env.example .env
    ```

    Sau đó cập nhật các giá trị trong file `.env` với thông tin thực tế của bạn.

3. Chạy lệnh `docker-compose up --build -d`

    ```bash
    docker-compose up --build -d
    ```

4. Truy cập vào `http://localhost` để sử dụng và `http://localhost:8080` sử dụng API Gateway

5. Đăng ký tài khoản và sử dụng

6. Chúc bạn một ngày tốt lành!

## Lưu ý bảo mật

-   KHÔNG COMMIT file `.env` lên git repository
-   Giữ các thông tin credentials an toàn và không chia sẻ
-   Thay đổi các credentials định kỳ để đảm bảo an toàn
