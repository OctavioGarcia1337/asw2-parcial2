FROM node:16

ADD . /frontend
WORKDIR /frontend

COPY package.json ./
COPY package-lock.json ./
COPY ./public/images ./public/images
COPY ./ ./
RUN npm i
CMD ["npm", "run", "start"]
