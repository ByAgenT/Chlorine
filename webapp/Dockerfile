# Build web application
FROM node:13.11 as build-stage

WORKDIR /webapp

COPY package*.json /webapp/
RUN npm install
COPY . /webapp/
RUN npm run build

# Use nginx to serve web application
FROM nginx:1.17

COPY --from=build-stage /webapp/dist/ /etc/nginx/html/
COPY config/nginx.conf /etc/nginx/nginx.conf
