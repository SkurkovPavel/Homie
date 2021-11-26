CREATE TABLE scenarios (
                           id bigserial not null primary key,
                           scenario_id varchar not null unique,
                           name varchar
);