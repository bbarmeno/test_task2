-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.user_account
(
    "ID" serial,
    "DATE_CREATED" timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    "BALANCE" float DEFAULT 0,
    PRIMARY KEY ("ID")
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.user_account
    OWNER to root;

INSERT INTO public.user_account("DATE_CREATED", "BALANCE")
	VALUES(CURRENT_TIMESTAMP, '10000.00');

INSERT INTO public.user_account("DATE_CREATED", "BALANCE")
	VALUES(CURRENT_TIMESTAMP, '0.00');

INSERT INTO public.user_account("DATE_CREATED", "BALANCE")
	VALUES(CURRENT_TIMESTAMP, '500.00');

CREATE TABLE public.transactions
(
    "ID" VARCHAR(52),
    "DATE_CREATED" timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    "AMOUNT" float,
    "STATE" VARCHAR(4),
    "USER_ID" integer,
    PRIMARY KEY ("ID"),
    UNIQUE ("ID"),
    CONSTRAINT "USER_ID" FOREIGN KEY ("USER_ID")
        REFERENCES public.user_account ("ID") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
);

ALTER TABLE IF EXISTS public.transactions
    OWNER to root;

INSERT INTO public.transactions("ID", "DATE_CREATED", "AMOUNT", "STATE", "USER_ID")
	VALUES('a',CURRENT_TIMESTAMP, '10000.00', 'win', '1');

INSERT INTO public.transactions("ID", "DATE_CREATED", "AMOUNT", "STATE", "USER_ID")
	VALUES('abc', CURRENT_TIMESTAMP, '500.00', 'win', '3');
    
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
