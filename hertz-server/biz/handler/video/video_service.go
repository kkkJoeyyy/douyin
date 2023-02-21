// Code generated by hertz generator.

package video

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	video "github.com/simplecolding/douyin/hertz-server/biz/model/hertz/video"
	"github.com/simplecolding/douyin/hertz-server/biz/mw"
	"github.com/simplecolding/douyin/hertz-server/biz/orm/dal"
	"github.com/simplecolding/douyin/hertz-server/biz/orm/model"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// VideoPublish .
// @router /douyin/publish/action [POST]
func VideoPublish(ctx context.Context, c *app.RequestContext) {
	println("fsdasfdd")
	//req.Data type : bytes
	var err error
	var req video.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		println("asdffadasdffasd",err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	//req.Token, err = c.Get("token")
	resp := new(video.DouyinPublishActionResponse)
	// jwt授权，从token中获取uid
	u, _ := c.Get(mw.IdentityKey)
	println("video:", u.(*mw.Claim).ID, u.(*mw.Claim).Username)
	//tk := req.Token
	//jwtToken, err := mw.JwtMiddleware.ParseTokenString(tk)
	//tmp := jwt.ExtractClaimsFromToken(jwtToken)
	// 检查文件类型
	fileType := http.DetectContentType(req.Data)
	userName := u.(*mw.Claim).Username
	if len(userName) == 0 {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + userName
	//fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		println("get filetype failed")
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	if fileType != "video/mp4" {
		resp.StatusCode = int32(1)
		resp.StatusMsg = "video type incorrect"
		c.JSON(consts.StatusOK, resp)
		println("video incorrect")
		return
	}
	filePath := filepath.Join("./biz/public", fileName+".mp4")
	newFile, err := os.Create(filePath)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		println("create file failed")
		return
	}
	defer newFile.Close()
	if _, err := newFile.Write(req.Data); err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		println("write file failed")
		return
	}

	playUrl := "http://127.0.0.1/public" + fileName + ".mp4"
	err = dal.Video.WithContext(ctx).Create(&model.Video{UID: u.(*mw.Claim).ID, PlayURL: playUrl, CoverURL: "coverUrl"})
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		println("write to database failed")
		return
	}
	println(filePath)
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	println(len(req.Data))
	c.JSON(consts.StatusOK, resp)
}

// GetPublishList .
// @router /douyin/publish/list [GET]
func GetPublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// jwt授权，从token中获取uid
	u, _ := c.Get(mw.JwtMiddleware.IdentityKey)
	println("video list:", u.(*mw.Claim).ID, u.(*mw.Claim).Username)

	resp := new(video.DouyinFavoriteListResponse)

	data, err := dal.Video.Where(dal.Video.UID.Eq(u.(*mw.Claim).ID)).Find()
	if err != nil {
		println("query database failed")
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	var v []*video.Video
	for _,d  := range data {
		var tmp video.Video
		tmp.Id = d.Vid
		tmp.CoverUrl = d.CoverURL
		tmp.PlayUrl = d.PlayURL
		// 还有很多没有实现的
		v = append(v, &tmp)
	}
	resp.VideoList = v
	c.JSON(consts.StatusOK, resp)
}

// GetFeed .
// @router /douyin/feed [GET]
func GetFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.DouyinFeedResponse)
	// todo
	// 30 videos for a single time
	limit := 30 
	// data, err := dal.Video.Order("created_at desc").Limit(limit).Find()
	data, err := dal.Video.Limit(limit).Find()
	if err != nil {
		println("query database failed")
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	var v []*video.Video
	for _,d  := range data {
		var tmp video.Video
		tmp.Id = d.Vid
		tmp.CoverUrl = d.CoverURL
		tmp.PlayUrl = d.PlayURL
		v = append(v, &tmp)
	}
	resp.VideoList = v
	c.JSON(consts.StatusOK, resp)
}
