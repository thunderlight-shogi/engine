FROM node:21-alpine3.19
COPY website /app/website
WORKDIR /app/website
RUN npm install && npm run build
