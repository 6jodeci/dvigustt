CREATE TABLE requests (
  id SERIAL PRIMARY KEY,
  ip VARCHAR(50) NOT NULL,
  datetimez timestamp NOT NULL
);