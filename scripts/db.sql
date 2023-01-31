/* DROP TABLE IF EXISTS users; */
/* DROP TABLE IF EXISTS positions; */
/* DROP VIEW IF EXISTS all_positions; */
/* DROP VIEW IF EXISTS usr; */

CREATE TABLE users (
  user_id int GENERATED ALWAYS AS IDENTITY,
  email varchar(320) NOT NULL UNIQUE,
  password varchar NOT NULL,

  CONSTRAINT pk_users_user_id PRIMARY KEY(user_id)
);

CREATE TABLE positions (
  position_id int GENERATED ALWAYS AS IDENTITY,
  open_date date NOT NULL DEFAULT CURRENT_DATE,
  pair varchar(12) NOT NULL,
  reason text NOT NULL,
  according_to_plan bool NOT NULL DEFAULT true,
  percentage_risk int NOT NULL,
  direction varchar(5) NOT NULL,
  deposit decimal NOT NULL,
  open_price decimal NOT NULL,
  stop_loss_price decimal NOT NULL,
  take_profit_price decimal NOT NULL,
  close_price decimal,
  user_id int NOT NULL,

  CONSTRAINT pk_positions_position_id PRIMARY KEY(position_id),
  CONSTRAINT fk_positions_user_user_id FOREIGN KEY(user_id) REFERENCES users(user_id)
);

CREATE VIEW all_positions AS
  SELECT position_id,
    open_date,
    pair,
    reason,
    according_to_plan,
    percentage_risk,
    direction,
    deposit,
    open_price,
    stop_loss_price,
    take_profit_price,
    close_price,
    user_id
  FROM positions
  WITH LOCAL CHECK OPTION;

CREATE VIEW usr AS
  SELECT user_id, email, password
  FROM users
  WITH LOCAL CHECK OPTION;

/* INSERT INTO users(email, password) */
/* VALUES */
/* ('google@gmail.com', 'password1'), */
/* ('yandex@yandex.ru', 'password2'), */
/* ('apple@icloud.com', 'password3'); */

/* INSERT INTO positions(pair, reason, percentage_risk, direction, deposit, open_price, stop_loss_price, take_profit_price, close_price, user_id) */
/* VALUES */
/* ('btc/usdt', 'some reason1', 1, 'long', 100, 20000, 19000, 23000, null, 1), */
/* ('btc/usdt', 'some reason2', 3, 'short', 300, 12000, 13000, 10000, 10000, 3), */
/* ('btc/usdt', 'some reason3', 1, 'long', 100, 20000, 19000, 23000, null, 1), */
/* ('btc/usdt', 'some reason4', 1, 'short', 200, 15000, 16000, 11000, 11000, 1), */
/* ('btc/usdt', 'some reason5', 1, 'long', 500, 22000, 21000, 25000, null, 1); */
