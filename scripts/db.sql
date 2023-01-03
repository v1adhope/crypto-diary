CREATE TABLE IF NOT EXISTS users (
  user_id int GENERATED ALWAYS AS IDENTITY,
  email varchar(320) NOT NULL UNIQUE,
  salt varchar NOT NULL,
  passhash varchar NOT NULL,

  CONSTRAINT pk_users_user_id PRIMARY KEY(user_id)
);

CREATE TABLE IF NOT EXISTS positions (
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

  CONSTRAINT pk_positions_position_id PRIMARY KEY(position_id),

  CONSTRAINT chk_positions_percentage_risk CHECK(percentage_risk > 0 AND percentage_risk < 100),
  CONSTRAINT chk_positions_direction CHECK(direction = 'long' OR direction = 'short'),
  CONSTRAINT chk_positions_deposit CHECK(deposit > 0),
  CONSTRAINT chk_positions_open_price CHECK(open_price > 0),
  CONSTRAINT chk_positions_stop_loss_price CHECK(stop_loss_price < open_price),
  CONSTRAINT chk_positions_take_profit_price CHECK(take_profit_price > open_price),
  CONSTRAINT chk_positions_close_price CHECK(close_price >= -deposit)
);
