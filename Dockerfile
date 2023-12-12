FROM node:lts as build

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY tailwind.config.js .

COPY tsconfig.json .

COPY src/ src/

RUN npm run tailwind

RUN	npm run build

FROM golang:bookworm as compile

WORKDIR /go/src/app

COPY --from=build app/dist/ dist/

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/base-debian12

COPY --from=build app/dist/ dist/

COPY --from=compile /go/bin/app /

COPY src/ src/

CMD ["./app"]
