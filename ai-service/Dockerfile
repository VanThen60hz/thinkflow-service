# Sử dụng hình ảnh Python chuẩn
FROM python:3.10-slim

# Đặt thư mục làm việc trong container
WORKDIR /app

# Copy file requirements trước để cache được bước cài thư viện
COPY requirements.txt .

# Cài đặt các thư viện cần thiết
RUN pip install --no-cache-dir -r requirements.txt

# Copy toàn bộ mã nguồn vào thư mục /app trong container
COPY . .

# Thiết lập biến môi trường để Flask chạy production
ENV PYTHONUNBUFFERED=True
ENV PORT=5000

# EXPOSE cổng mà app sẽ chạy
EXPOSE 5000

# Lệnh khởi động container
CMD ["gunicorn", "-b", "0.0.0.0:5000", "app:app"]
