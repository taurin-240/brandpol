FROM python:3.11
COPY app/ /app/
COPY requirements.txt /app/
RUN pip install --no-cache-dir --upgrade -r ./app/requirements.txt
WORKDIR /app
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1
CMD ["uvicorn", "main:app", "--host", "localhost", "--port", "8000"]