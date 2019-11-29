FROM node:10.17-alpine as builder

WORKDIR /app
COPY ./public ./public
COPY ./src ./src
COPY babel.config.js .
COPY package.json .
COPY vue.config.js .
COPY yarn.lock .
COPY .eslintrc.js .
RUN yarn install
RUN yarn build

# runtime image
FROM node:10.17-alpine

WORKDIR /app
COPY ./docker ./
COPY --from=builder /app/dist ./dist
RUN yarn install

EXPOSE 8080
CMD [ "node", "server.js" ]
