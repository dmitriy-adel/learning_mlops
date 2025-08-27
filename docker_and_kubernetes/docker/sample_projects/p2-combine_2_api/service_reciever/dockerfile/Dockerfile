FROM python:3.10

WORKDIR /app
COPY ./requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 8088
CMD ["uvicorn", "service_reciever.app.main:app_receiver", "--host", "0.0.0.0", "--port", "8088"]
