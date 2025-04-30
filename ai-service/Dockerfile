FROM python:3.10-slim

WORKDIR /app

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

ENV PYTHONUNBUFFERED=True
ENV PORT=5000

EXPOSE 5000

CMD ["gunicorn", "--timeout", "3600", "--bind", "0.0.0.0:5000", "app:app"]