FROM scratch

WORKDIR /app

COPY dist/authenticator /app/

ENTRYPOINT ["./authenticator"]
