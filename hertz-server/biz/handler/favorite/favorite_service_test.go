package favorite

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/simplecolding/douyin/hertz-server/biz/model/hertz/favorite"
	"reflect"
	"strings"
	"testing"
)

func TestFavoriteAction(t *testing.T) {
	h := server.Default()
	h.POST("/douyin/favorite/action", FavoriteAction)
	url := "/douyin/favorite/action"
	method := "POST"

	reqBody := strings.NewReader("action_type=1&token=FtWeIqOfPtokcIwGWNqAPKnQqhOfGZxfUuUcawOEiAVCQeCXzj&video_id=7")
	w := ut.PerformRequest(h.Engine,
		method,
		url,
		&ut.Body{reqBody,reqBody.Len()},
		ut.Header{"Content-Type", "application/x-www-form-urlencoded"},
	)
	resp := w.Result()
	//assert.DeepEqual(t, 200, resp.StatusCode())
	println(string(resp.Body()))
}

func TestFavoriteQueryUser(t *testing.T) {
	type args struct {
		uid int64
	}
	tests := []struct {
		name string
		args args
		want favorite.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FavoriteQueryUser(tt.args.uid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FavoriteQueryUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFavoriteList(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *app.RequestContext
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
