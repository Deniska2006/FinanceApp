CREATE TABLE IF NOT EXISTS public.costs (
    id integer PRIMARY KEY,
    name TEXT NOT NULL,
    price integer NOT NULL,
    createdtime  TIMESTAMPTZ
);