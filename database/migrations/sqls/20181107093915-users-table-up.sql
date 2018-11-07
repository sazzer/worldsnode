CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    version UUID NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    updated TIMESTAMP WITH TIME ZONE NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    email VARCHAR(260) NULL,
    providers JSONB NOT NULL
);
