CREATE TABLE telemetry (

                           id          uuid PRIMARY KEY ,
                           drone_id uuid NOT NULL REFERENCES drones(id),
                           timestamp   timestamptz NOT NULL,
                           latitude    REAL,
                           longitude   REAL,
                           altitude    REAL,
                           speed       REAL,
                           battery    REAL,
                           temperature REAL,
                           heading     REAL,
                           signal      REAL,
                           latency     INTEGER,
                           packet_loss  REAL

)
