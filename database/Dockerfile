FROM postgres:11.0-alpine

ENV PGDATA=/pgdata

# Install Node and NPM
RUN apk add --update nodejs nodejs-npm

# Install db-migrate
RUN npm install -g db-migrate db-migrate-pg

# Copy migration files
COPY docker/* /docker-entrypoint-initdb.d/
COPY migrations /migrations/migrations
COPY database.json /migrations

# Start Postgres
RUN sed -i.bak "s/listen_addresses=''/listen_addresses='127.0.0.1'/" /usr/local/bin/docker-entrypoint.sh
RUN /usr/local/bin/docker-entrypoint.sh postgres --version

FROM postgres:11.0-alpine
ENV PGDATA=/pgdata
COPY --from=0 /pgdata /pgdata
