FROM scratch

COPY public /app/
COPY cokut-server /app/

WORKDIR /app

CMD ["./cokut-server"]

EXPOSE 4000