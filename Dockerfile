FROM alpine
COPY ./yedit /app/yedit
RUN chmod +x /app/yedit