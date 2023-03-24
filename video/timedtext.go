package video

import (
	"net/http"

	"github.com/sujjadapp/videoproxy/db"
	"github.com/sujjadapp/videoproxy/request"
	"github.com/sujjadapp/videoproxy/util"
	"github.com/sujjadapp/youtubevideoparser"
)

// info 解析存在缓存,此处ProxyCall也缓存
func outPutTimedText(w http.ResponseWriter, r *http.Request, info *youtubevideoparser.VideoInfo) error {
	var (
		useLang = ""
		lang    = r.URL.Query().Get("lang")
		url     = ""
	)
	for _, item := range info.Captions {
		useLang = item.Language
		url = item.URL
		if lang == "" || useLang == lang {
			break
		}
	}
	if url == "" || useLang == "" {
		http.Error(w, "lang not found", http.StatusNotFound)
		return nil
	}
	var hook = func(data []byte, status int) {
		if status == http.StatusOK {
			if err := db.SaveCaption(info.ID, useLang, data); err != nil {
				util.Log.Print(err)
			}
		}
	}
	return request.ProxyCall(w, url, videoClient, r.Header, hook)
}
