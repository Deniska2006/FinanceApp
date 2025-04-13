CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY,
    mail text NOT NULL,
    hashed_password TEXT NOT NULL
);