FROM scratch

ENV PORT 8000
EXPOSE $PORT

COPY tutor /
CMD ["/tutor"]
