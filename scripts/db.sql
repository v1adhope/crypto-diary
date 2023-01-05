DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS positions;

CREATE TABLE users (
  user_id int GENERATED ALWAYS AS IDENTITY,
  email varchar(320) NOT NULL UNIQUE,
  password varchar NOT NULL,

  CONSTRAINT pk_users_user_id PRIMARY KEY(user_id)
);

CREATE TABLE positions (
  position_id int GENERATED ALWAYS AS IDENTITY,
  open_date date NOT NULL DEFAULT CURRENT_DATE,
  pair varchar(5) NOT NULL,
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

/*   CONSTRAINT chk_positions_percentage_risk CHECK(percentage_risk > 0 AND percentage_risk < 100), */
/*   CONSTRAINT chk_positions_direction CHECK(direction = 'long' OR direction = 'short'), */
/*   CONSTRAINT chk_positions_deposit CHECK(deposit > 0), */
/*   CONSTRAINT chk_positions_open_price CHECK(open_price > 0), */
/*   CONSTRAINT chk_positions_stop_loss_price CHECK(stop_loss_price < open_price), */
/*   CONSTRAINT chk_positions_take_profit_price CHECK(take_profit_price > open_price), */
/*   CONSTRAINT chk_positions_close_price CHECK(close_price >= -deposit) */
/* ); */

-- TODO: USE JWT
INSERT INTO users(email, password)
VALUES
('google@gmail.com', 'password1'),
('yandex@yandex.ru', 'password2'),
('apple@icloud.com', 'password3');

INSERT INTO positions(pair, reason, percentage_risk, direction, deposit, open_price, stop_loss_price, take_profit_price, close_price, user_id)
VALUES
('btc', 'some reason1', 1, 'long', 100, 20000, 19000, 23000, null, 1),
('btc', 'some reason2', 3, 'short', 300, 12000, 13000, 10000, 10000, 3);
