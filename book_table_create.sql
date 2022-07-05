-- Table: public.books

-- DROP TABLE IF EXISTS public.books;

CREATE TABLE IF NOT EXISTS public.books
(
    id serial NOT NULL,
    book_id character varying(150) COLLATE pg_catalog."default",
    book_name character varying(150) COLLATE pg_catalog."default",
    CONSTRAINT books_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.books
    OWNER to postgres;