FROM node:21-alpine3.19
COPY website /app/website
WORKDIR /app/website
# && npm run build
RUN npm install
CMD ["npm", "run", "dev"]
EXPOSE 5173
