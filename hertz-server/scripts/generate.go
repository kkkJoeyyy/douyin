package main

import (
	"github.com/simplecolding/douyin/hertz-server/biz/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

//// FavoriteQuerier Querier Dynamic SQL
//type FavoriteQuerier interface {
//	// FilterWithVidAndUid SELECT * FROM @@table WHERE vid = @vid{{if uid !=-1}} AND uid = @uid{{end}}
//	FilterWithVidAndUid(vid, uid int64) ([]gen.T, error)
//}
//
//// CommentQuerier UserQuerier Querier Dynamic SQL
//type CommentQuerier interface {
//
//}
//
//// VideoQuerier Querier Dynamic SQL
//type VideoQuerier interface {
//
//}
//
//// UserQuerier VideoQuerier Querier Dynamic SQL
//type UserQuerier interface {
//
//}
func main() {
	gormdb, _ := gorm.Open(mysql.Open(utils.MySQLDefaultDSN))
	g := gen.NewGenerator(gen.Config{
		OutPath: "../biz/orm/dal",
		ModelPkgPath: "../biz/orm/model",
		Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
	})
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following convention
	g.ApplyBasic(g.GenerateAllTable()...) //
	//g.ApplyBasic(model.UserAuth{},model.Comment{},model.Video{})
	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(FavoriteQuerier) {}, model.Favorite{})
	//g.ApplyInterface(func(UserQuerier) {}, model.UserAuth{})
	//g.ApplyInterface(func(CommentQuerier) {}, model.Comment{})
	//g.ApplyInterface(func(VideoQuerier) {}, model.Video{})

	// Generate the code
	g.Execute()
}
