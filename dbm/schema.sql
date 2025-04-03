CREATE TABLE IF NOT EXISTS public.costs (
    id integer PRIMARY KEY,
    name TEXT NOT NULL,
    price DECIMAL NOT NULL,
    createdtime  TIMESTAMPTZ
);