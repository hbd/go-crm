CREATE SEQUENCE unique_id_sequence
    START WITH 1
    INCREMENT BY 1
    MINVALUE 1
    NO MAXVALUE
    CACHE 1;

CREATE TABLE client(
    name VARCHAR(255),
    id VARCHAR(18) UNIQUE,
    contact TEXT,
    stage VARCHAR(18),
    description TEXT
);

CREATE TABLE stage(
    stageType VARCHAR(20),
    id VARCHAR(18) UNIQUE,
    startDate TIMESTAMP,
    endDate TIMESTAMP,
    stageNotes TEXT
);