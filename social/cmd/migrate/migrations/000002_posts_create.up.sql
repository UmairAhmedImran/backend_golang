CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY NOT NULL,
    title text NOT NULL,
    user_id bigint NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP(0) with time zone NOT NULL default NOW()
);