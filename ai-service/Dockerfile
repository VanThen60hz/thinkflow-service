FROM python:3.10-slim

WORKDIR /app

RUN apt-get update && apt-get install -y build-essential gcc

COPY requirements.txt .

RUN pip install --upgrade pip && \
    pip install --cache-dir=/tmp/pip-cache -r requirements.txt

COPY . .

RUN apt-get remove -y build-essential gcc && apt-get autoremove -y && rm -rf /var/lib/apt/lists/*

ENV PYTHONUNBUFFERED=True
ENV PORT=5000

EXPOSE 5000

CMD ["gunicorn", "--timeout", "3600", "--bind", "0.0.0.0:5000", "app:app"]