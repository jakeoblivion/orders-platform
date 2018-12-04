# 1. FRONT-END INTERFACE
FROM node:10.8.0-alpine as front-end-builder

WORKDIR /build
COPY package.json package.json
COPY package-lock.json package-lock.json
COPY .eslintrc.js .eslintrc.js
COPY src src
COPY public public

RUN npm install --production --no-optional
RUN npm run build

# 2. BACK-END SERVER
FROM golang:alpine as back-end-builder

WORKDIR /build
COPY backend/src/orders-management backend/src/orders-management

RUN apk add git
RUN go get github.com/mrjones/oauth
RUN GOOS=linux GOARCH=amd64 go build -o main backend/src/orders-management/*.go

# 3. START SERVER AND SERVE FRONT-END
FROM alpine:3.8
WORKDIR /app
RUN apk add git

COPY --from=front-end-builder /build/dist /app/dist
COPY --from=back-end-builder /build/main /app

RUN addgroup -S main_user
RUN adduser -S -G main_user main_user
USER main_user

CMD ["./main"]

