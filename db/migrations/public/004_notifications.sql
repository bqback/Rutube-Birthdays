CREATE TABLE IF NOT EXISTS public.notifications
(
    id_source int NOT NULL,
    id_subscriber int NOT NULL,
    CONSTRAINT notifications_pkey PRIMARY KEY (id_source, id_subscriber),
    CONSTRAINT source_id_user_fkey FOREIGN KEY (id_source)
        REFERENCES public.user (id_user) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT subscriber_id_user_fkey FOREIGN KEY (id_subscriber)
        REFERENCES public.user (id_user) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID    
);

---- create above / drop below ----

DROP TABLE IF EXISTS public.notifications;