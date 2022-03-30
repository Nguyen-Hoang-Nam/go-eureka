FROM node:16-alpine3.14
WORKDIR /app
ENV NODE_ENV=production
COPY wait-for wait-for
COPY ["package.json", "package-lock.json*", "./"]
RUN npm install --production
COPY . .
ENTRYPOINT [ "./wait-for", "eureka:9000", "--", "node", "app.js" ]
