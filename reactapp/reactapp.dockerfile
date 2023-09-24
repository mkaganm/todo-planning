# Build Stage
FROM node:16 as build-stage

WORKDIR /app

COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Production Stage
FROM node:16

WORKDIR /app

COPY --from=build-stage /app/build/ ./build/

RUN npm install -g serve

CMD ["serve", "-s", "build", "-l", "8080"]