CREATE DATABASE plantinfo;

\c plantinfo

CREATE TABLE IF NOT EXISTS plants (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    common_name text NOT NULL,
    seed_company text NOT NULL,
    expected_days_to_harvest integer NOT NULL,
    type text NOT NULL,
    ph_low real NOT NULL,
    ph_high real NOT NULL,
    ec_low real NOT NULL,
    ec_high real NOT NULL,
    version integer NOT NULL DEFAULT 1
);

GRANT SELECT, INSERT, UPDATE, DELETE ON plants TO tester;
GRANT USAGE, SELECT ON SEQUENCE plants_id_seq TO tester;