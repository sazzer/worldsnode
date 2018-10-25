FROM node:10.12.0-alpine

ENV PORT=3000
ENV NODE_ENV=production

# Copy application into container
COPY . /opt/worlds/

# Download dependencies
WORKDIR /opt/worlds
RUN yarn install

# Set the startup command
CMD ./node_modules/.bin/babel-node src/index.js

# Runs on port 3000 at present
EXPOSE 3000
