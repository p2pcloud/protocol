FROM node:16 as builder

WORKDIR /app

COPY package* /app/
RUN npm ci

COPY ./contracts /app/contracts
COPY ./test /app/test
COPY ./hardhat.config.ts /app/

RUN npm run test
RUN npm run compile

RUN npm run flat

# main container

FROM halverneus/static-file-server:v1.8.9

WORKDIR /web/

COPY --from=builder /app/artifacts ./artifacts
COPY --from=builder /app/typechain-types ./typechain-types
