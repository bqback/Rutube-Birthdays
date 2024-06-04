CREATE TABLE IF NOT EXISTS public.user
(
    id_user int UNIQUE NOT NULL,
    name text NOT NULL DEFAULT '',
    surname text NOT NULL DEFAULT '',
    email text NOT NULL,
    dob date NOT NULL,
    CONSTRAINT user_id_auth_fkey FOREIGN KEY (id_user)
        REFERENCES public.auth (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
);

---- create above / drop below ----

DROP TABLE IF EXISTS public.user;
