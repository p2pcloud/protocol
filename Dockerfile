FROM node:18-alpine

# Create app directory
WORKDIR /app

COPY package* ./
RUN npm ci

COPY . /app/
