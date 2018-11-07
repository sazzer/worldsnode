CREATE TABLE oauth2_clients(
    client_id UUID PRIMARY KEY,
    version UUID NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    updated TIMESTAMP WITH TIME ZONE NOT NULL,
    owner_id UUID NOT NULL REFERENCES users(user_id),
    client_name TEXT NOT NULL,
    client_secret UUID NOT NULL,
    redirect_uris TEXT[] NOT NULL,
    response_types TEXT[] NOT NULL,
    grant_types TEXT[] NOT NULL
);
