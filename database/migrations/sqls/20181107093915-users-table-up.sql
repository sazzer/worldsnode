CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    version UUID NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    updated TIMESTAMP WITH TIME ZONE NOT NULL,
    display_name TEXT NOT NULL,
    email TEXT NULL,
    providers JSONB NOT NULL
);
