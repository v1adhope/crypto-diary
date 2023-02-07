package v1

// TODO
// func (r *Deps) AuthorizeJWT() gin.HandlerFunc {
// 	const _bearerSchema = "Bearer "
//
// 	return func(c *gin.Context) {
// 		header := c.GetHeader("Autorization")
// 		if header == "" {
// 			errorResponse(c, http.StatusUnauthorized, "empty auth header")
// 			return
// 		}
//
// 		clientToken := header[len(_bearerSchema):]
//
// 		_, err := r.Auth.ValidateToken(clientToken)
// 		if err != nil {
// 			r.Logger.Logger.Error().Err(err).Msg("AuthMiddleware")
// 			errorResponse(c, http.StatusUnauthorized, "invalid token")
// 			return
// 		}
// 	}
// }
