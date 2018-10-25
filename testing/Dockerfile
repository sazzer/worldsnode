FROM node:10.12.0-alpine

ENV SERVICE_URI=http://service:3000

# Set up Dockerize
RUN apk add --no-cache openssl

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
    
# Copy application into container
COPY . /opt/worlds/

# Download dependencies
WORKDIR /opt/worlds
RUN yarn install

# Set the startup command
CMD dockerize -timeout 30s -wait tcp://database:5432 -wait tcp://service:3000 yarn test-suite
