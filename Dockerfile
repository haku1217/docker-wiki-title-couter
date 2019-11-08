FROM alpine:3.7

WORKDIR /app
COPY ./app /app

RUN apk add python3 wget gzip && rm -rf /var/cache/apk/*
RUN wget https://dumps.wikimedia.org/jawiki/latest/jawiki-latest-all-titles.gz && gzip -d jawiki-latest-all-titles.gz

ENTRYPOINT ["python3", "/app/wikicount.py"]