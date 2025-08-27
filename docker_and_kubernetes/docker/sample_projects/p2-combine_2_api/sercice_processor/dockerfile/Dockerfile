FROM python:3.10

WORKDIR /app
COPY ./requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 8087
CMD ["uvicorn", "sercice_processor.app.main:app_processor", "--host", "0.0.0.0", "--port", "8087"]
