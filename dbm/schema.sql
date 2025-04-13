CREATE TABLE IF NOT EXISTS public.costs (
    id           SERIAL PRIMARY KEY,
    uid          INTEGER NOT NULL REFERENCES users(id),
    name         TEXT NOT NULL,
    price        DECIMAL NOT NULL,
    createdtime  TIMESTAMPTZ
);