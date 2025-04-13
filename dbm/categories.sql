CREATE TABLE IF NOT EXISTS public.categories (
    id SERIAL PRIMARY KEY,
    uid  INTEGER NOT NULL REFERENCES users(id),
    name TEXT NOT NULL
);