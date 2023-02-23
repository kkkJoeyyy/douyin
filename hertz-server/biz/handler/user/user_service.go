// Code generated by hertz generator.

package user

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/simplecolding/douyin/hertz-server/biz/redis"
	"github.com/simplecolding/douyin/hertz-server/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "github.com/simplecolding/douyin/hertz-server/biz/model/hertz/user"
	_ "github.com/simplecolding/douyin/hertz-server/biz/mw"
	_ "github.com/simplecolding/douyin/hertz-server/biz/orm" //
	"github.com/simplecolding/douyin/hertz-server/biz/orm/dal"
	"github.com/simplecolding/douyin/hertz-server/biz/orm/model"
)

// UserLogin .
// @router /douyin/user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	u, err := dal.UserAuth.Where(dal.UserAuth.UserName.Eq(req.Username)).First()
	if err != nil {
		c.JSON(consts.StatusBadRequest, "用户不存在!!!")
	}
	if u.Password != req.Password {
		c.JSON(consts.StatusBadRequest, "密码错误!!!")
	}
	redisCtx, cancel := context.WithTimeout(context.Background(), 500*time.Hour)
	defer cancel()
	resp := new(user.DouyinUserLoginResponse)

	tk := utils.GenToken()
	println("token", tk)
	resp.Token = tk
	redis.RD.Set(redisCtx, tk, req.Username, 500*time.Hour)
	redis.RD.Set(redisCtx, req.Username, u.UID, 500*time.Hour)

	resp.StatusCode = 0
	resp.StatusMsg = "success"
	resp.UserId = u.UID
	c.JSON(consts.StatusOK, resp)
}

// UserRegister .
// @router /douyin/user/register [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRegisterRequest
	resp := new(user.DouyinUserRegisterResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	username := req.Username
	password := req.Password
	if len(username) < 6 || len(password) < 6 {
		resp.StatusCode = 1
		resp.StatusMsg = "username must gt 6 and password must gt 6!!!"
		c.JSON(consts.StatusOK, resp)
		return
	}
	// query
	u, err := dal.UserAuth.Where(dal.UserAuth.UserName.Eq(username)).First()
	if u != nil {
		resp.UserId = -1
		resp.StatusCode = 1
		resp.StatusMsg = "failed!  username已存在！！！" + username
		c.JSON(consts.StatusOK, resp)
		return
	}

	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	t := time.Now().In(cstSh)
	userinfo := model.UserAuth{UserName: username, Password: password, CreatedAt: t}
	err = dal.UserAuth.Create(&userinfo)
	if err != nil {
		resp.UserId = -1
		resp.StatusCode = 1
		resp.StatusMsg = "failed!  插入失败！！！"
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.UserId = userinfo.UID
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	c.JSON(0, resp)
}

// UserInfo .
// @router /douyin/user [POST]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// 鉴权
	flag, _, uid := utils.Auth(ctx, req.Token)
	if !flag {
		c.JSON(consts.StatusBadRequest, "token错误")
		return
	}
	dal.UserAuth.Where(dal.UserAuth.UID.Eq(uid))

	totalFavorited := int64(0)
	v, err := dal.Video.Where(dal.Video.UID.Eq(uid)).Find()
	for _, t := range v {
		tmpcount, _ := dal.Favorite.Where(dal.Favorite.Vid.Eq(t.Vid)).Count()
		totalFavorited += tmpcount
	}
	// 低性能代码
	workCount, _ := dal.Video.Where(dal.Video.UID.Eq(uid)).Count()
	favoriteCount, _ := dal.Favorite.Where(dal.Favorite.UID.Eq(uid)).Count()
	userInfoDB, err := dal.UserAuth.Where(dal.UserAuth.UID.Eq(uid)).First()
	if err != nil {
		println("database err")
	}
	userInfoDB.WorkCount = workCount
	// 喜欢数
	userInfoDB.FavoriteCount = favoriteCount
	// 获得赞数量
	userInfoDB.TotalFavorite = strconv.FormatInt(totalFavorited, 10)
	dal.UserAuth.Save(userInfoDB)

	resp := new(user.DouyinUserResponse)
	resp.StatusCode = int32(0)
	resp.StatusMsg = "success"
	resp.User = &user.User{
		Id:              userInfoDB.UID,
		Name:            userInfoDB.UserName,
		FollowCount:     userInfoDB.FollowCount,
		IsFollow:        userInfoDB.IsFollow,
		Avatar:          userInfoDB.Avatar,
		BackgroundImage: userInfoDB.BackgroundImage,
		Signature:       userInfoDB.Signature,
		TotalFavorited:  userInfoDB.TotalFavorite,
		WorkCount:       userInfoDB.WorkCount,
		FavoriteCount:   userInfoDB.FavoriteCount,
	}
	c.JSON(http.StatusOK, resp)
	return
}

// QueryUser query userinfo
func QueryUser(uid int64) user.User {
	dal.UserAuth.Where(dal.UserAuth.UID.Eq(uid))
	totalFavorited := int64(0)
	v, err := dal.Video.Where(dal.Video.UID.Eq(uid)).Find()
	for _, t := range v {
		tmpcount, _ := dal.Favorite.Where(dal.Favorite.Vid.Eq(t.Vid)).Count()
		totalFavorited += tmpcount
	}
	// 低性能代码
	workCount, _ := dal.Video.Where(dal.Video.UID.Eq(uid)).Count()
	favoriteCount, _ := dal.Favorite.Where(dal.Favorite.UID.Eq(uid)).Count()
	userInfoDB, err := dal.UserAuth.Where(dal.UserAuth.UID.Eq(uid)).First()
	if err != nil {
		println("database err")
	}
	userInfoDB.WorkCount = workCount
	userInfoDB.FavoriteCount = favoriteCount
	userInfoDB.TotalFavorite = strconv.FormatInt(totalFavorited, 10)
	dal.UserAuth.Save(userInfoDB)

	return user.User{
		Id:              userInfoDB.UID,
		Name:            userInfoDB.UserName,
		FollowCount:     userInfoDB.FollowCount,
		IsFollow:        userInfoDB.IsFollow,
		Avatar:          userInfoDB.Avatar,
		BackgroundImage: userInfoDB.BackgroundImage,
		Signature:       userInfoDB.Signature,
		TotalFavorited:  userInfoDB.TotalFavorite,
		WorkCount:       userInfoDB.WorkCount,
		FavoriteCount:   userInfoDB.FavoriteCount,
	}
}
