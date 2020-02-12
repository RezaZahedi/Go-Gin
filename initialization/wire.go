package initialization
//
//import (
//	"github.com/RezaZahedi/Go-Gin/REST_api"
//	"github.com/RezaZahedi/Go-Gin/database"
//	"github.com/RezaZahedi/Go-Gin/product"
//	"github.com/RezaZahedi/Go-Gin/user"
//	"github.com/gin-gonic/gin"
//	"github.com/google/wire"
//)
//
//func initProductAPI(db *database.BookDB) *REST_api.ProductAPI {
//	wire.Build(product.ProvideProductRepository, product.ProvideProductService, REST_api.ProvideProductAPI)
//
//	return &REST_api.ProductAPI{}
//}
//func initUserAPI(db *database.UserDB) *REST_api.UserAPI {
//	wire.Build(user.ProvideUserRepository, user.ProvideUserService, REST_api.ProvideUserAPI)
//
//	return &REST_api.UserAPI{}
//}
//
//var ProductAPISet = wire.NewSet(
//	database.NewBookDB,
//	product.ProvideProductRepository,
//	product.ProvideProductService,
//	REST_api.ProvideProductAPI,
//)
//
//var UserAPISet = wire.NewSet(
//	database.NewUserDB,
//	user.ProvideUserRepository,
//	user.ProvideUserService,
//	REST_api.ProvideUserAPI,
//)
//
//func initUserBookREST (router *gin.Engine) error {
//	wire.Build(ProductAPISet, UserAPISet, REST_api.InitializeRoutes)
//
//	return nil
//}