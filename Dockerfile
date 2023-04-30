FROM node:18-alpine as builder

WORKDIR /app

COPY package* ./
RUN npm ci

COPY ./contracts ./contracts
COPY ./test ./test
COPY ./hardhat.config.ts ./

RUN npm run test
RUN npm run compile

RUN ls -ltha

RUN npm run flat

# main container

FROM halverneus/static-file-server:v1.8.9

WORKDIR /web/

COPY --from=builder /app/artifacts ./artifacts
COPY --from=builder /app/typechain-types ./typechain-types
