CREATE TABLE  todo
(
    id UUID PRIMARY KEY  DEFAULT gen_random_uuid(),
    title  text,
    body  text
)