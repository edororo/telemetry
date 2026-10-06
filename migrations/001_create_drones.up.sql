CREATE TABLE drones (
                        id       uuid PRIMARY KEY,
                        name text,
                        status   text,
                        created_at timestamptz,
                        last_seen  timestamptz
);

