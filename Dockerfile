FROM golang:1.9

ENV PORT 8000
EXPOSE $PORT
WORKDIR /app
COPY tutor /app
CMD ["/app/tutor"]
