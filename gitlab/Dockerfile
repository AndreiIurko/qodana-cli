FROM docker:27.5.1
WORKDIR /app
COPY dist/* /app/
COPY qodana /app/qodana-cli/
ENV PATH="$PATH:/app/qodana-cli"
RUN apk add --no-cache nodejs npm git bash