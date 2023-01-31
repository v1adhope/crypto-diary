package repository

// NOTE: Position
// p := &entity.Position{
// 	OpenDate:        "2023.01.17",
// 	Pair:            "btc/usdt",
// 	Risk:            "1",
// 	Reason:          "Some reason",
// 	AccordingToPlan: "true",
// 	Direction:       "short",
// 	Deposit:         "100",
// 	OpenPrice:       "20000",
// 	StopLossPrice:   "19000",
// 	TakeProfitPrice: "23000",
// 	ClosePrice:      "23000",
// 	UserID:          "1",
// }
// p.ValidPosition()
// if err != nil {
// 	logger.Fatal().Err(err).Send()
// }
// err = positionRepo.Create(context.Background(), p)
// if err != nil {
// 	logger.Fatal().Err(err).Send()
// }
// logger.Info().Msgf("", p)
//
// id := "3"
// err = positionRepo.Delete(context.Background(), &id)
// if err != nil {
// 	logger.Fatal().Err(err).Send()
// }
//
// id = "1"
// positions := make([]entity.Position, 0)
// positions, err = positionRepo.FindAll(context.Background(), &id)
// if err != nil {
// 	logger.Fatal().Err(err).Send()
// }
// logger.Info().Msgf("", positions)
//
// //NOTE: User
// userRepo := repository.NewUser(pgClient)
//
// u := &entity.User{
// 	Email:    "custom@custom.cu",
// 	Password: "password",
// }
// err = userRepo.CreateUser(context.Background(), u)
// if err != nil {
// 	logger.Fatal().Err(err).Send()
//
// }
//
// u2 := &entity.User{}
// email := "google@gmail.com"
// passwd := "password1"
// u2, err = userRepo.GetUser(context.Background(), &email, &passwd)
// if err != nil {
// 	logger.Fatal().Err(err).Send()
// }
// logger.Info().Msgf("", u2)
