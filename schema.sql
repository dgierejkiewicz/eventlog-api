CREATE TABLE public.event_log
(
    ts      bigint not null,
    type    varchar,
    content varchar
);

SELECT create_hypertable('event_log', 'ts', chunk_time_interval => 1209600);

INSERT INTO public.event_log (ts, type, content) VALUES('1622575864'::bigint, 'user_registration', 'Test content for user_registration');
INSERT INTO public.event_log (ts, type, content) VALUES('1622575865'::bigint, 'user_registration', 'Test content for user_registration');
INSERT INTO public.event_log (ts, type, content) VALUES('1622575866'::bigint, 'user_click', 'Test content for user_click action');
