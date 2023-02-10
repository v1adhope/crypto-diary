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
  strategically bool NOT NULL DEFAULT true,
  percentage_risk double precision NOT NULL,
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

CREATE VIEW get_all_positions AS
  SELECT position_id,
    open_date,
    pair,
    reason,
    strategically,
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

CREATE VIEW get_user AS
  SELECT user_id, email, password
  FROM users
  WITH LOCAL CHECK OPTION;
