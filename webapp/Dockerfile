FROM node:11.12-alpine

WORKDIR /usr/src/app
ADD . ./

ENV PATH /usr/src/app/node_modules/.bin:$PATH

RUN npm install

CMD ["npm", "start"]