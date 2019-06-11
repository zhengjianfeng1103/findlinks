FROM iron/base
WORKDIR /app
COPY findlinks /app/
ENTRYPOINT ["./findlinks"]