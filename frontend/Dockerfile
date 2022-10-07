FROM node:16-alpine

WORKDIR /app

COPY package.json .

RUN npm install

RUN npm i @vue/cli-service

COPY . .

EXPOSE 8081

CMD ["npm", "run", "serve"]