package views

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/go-xweb/uuid"
	"github.com/shumipro/meetapp/server/models"
	"github.com/shumipro/meetapp/server/oauth"
	"golang.org/x/net/context"
)

func readBodyAppInfo(body io.ReadCloser) (models.AppInfo, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return models.AppInfo{}, err
	}
	fmt.Println(string(data))

	var regAppInfo models.AppInfo
	if err := json.Unmarshal(data, &regAppInfo); err != nil {
		return models.AppInfo{}, err
	}

	return regAppInfo, nil
}

// 登録用にappInfoを加工して返します
func convertRegisterAppInfo(ctx context.Context, appInfo models.AppInfo) models.AppInfo {
	// アプリID（UUID）
	appInfo.ID = uuid.NewRandom().String()

	// 登録時刻、更新時刻
	nowTime := time.Now()
	appInfo.CreateAt = nowTime
	appInfo.UpdateAt = nowTime

	// 管理者設定
	a, _ := oauth.FromContext(ctx)
	for idx, m := range appInfo.Members {
		if m.UserID != a.UserID {
			continue
		}
		appInfo.Members[idx].IsAdmin = true
	}

	// メインの画像を設定
	appInfo.MainImage = appInfo.FirstImageURL()

	return appInfo
}