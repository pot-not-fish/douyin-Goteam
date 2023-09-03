package video

import (
	"douyin/pkg"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"douyin/apps/api/internal/logic/video"
	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
)

func dbcreate(db *gorm.DB, userinfo *pkg.User) {
	err := db.Model(&userinfo).Where("ID = ?", userinfo.ID).Update("WorkCount", userinfo.WorkCount+1).Error
	if err != nil {
		fmt.Printf("user updating error")
		return
	}
}

func PublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoPublishReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// token鉴权
		id, err := pkg.AuthToken(req.Token)
		if err != nil {
			return
		}

		// 数据库操作  储存命名为 用户id+标题+时间戳+文件格式
		db := svcCtx.DbEngin
		id_str := strconv.FormatInt(id, 10)
		t := strconv.FormatInt(time.Now().Unix(), 10)
		play := id_str + "_" + req.Title + "_" + t + ".mp4"
		fornowplay := id_str + "_" + t + "_" + ".mp4"
		cover := id_str + "_" + req.Title + "_" + t + ".jpg"

		// 缓存连接  将数据储存到list中
		// redisDb, err := pkg.RedisInit()
		// if err != nil {
		// 	fmt.Println("redis init error")
		// 	return
		// }
		redisDb := svcCtx.DbRedis

		// 查找userinfo信息
		userinfo := &pkg.User{}
		userinfo, err = pkg.RedisUserRead(db, redisDb, id)
		if err != nil {
			fmt.Println("redis search user error")
			return
		}

		// 数据库储存
		videoinfo := &pkg.Video{
			UserId:   id,
			PlayUrl:  play,
			CoverUrl: cover,
			Title:    req.Title,
		}
		db.Create(videoinfo)
		videodata, err := json.Marshal(videoinfo)
		if err != nil {
			fmt.Println("json error")
			return
		}

		// 设置时间为后面翻页推送设置使用
		llen, err := redisDb.LLen("videos").Result()
		if err != nil {
			fmt.Println("time error")
			return
		}
		if llen%10 == 0 {
			redisDb.LPush("time", videoinfo.CreatedAt)
		}
		redisDb.LPush("videos", string(videodata))

		// 获取文件数据 限制最大100MB
		err = r.ParseMultipartForm(100)
		if err != nil {
			return
		}
		file, _, err := r.FormFile("data")
		if err != nil {
			return
		}

		// 储存视频路径
		outplay, err := os.Create(`C:\Users\84023\Videos\` + fornowplay)
		if err != nil {
			return
		}

		// 储存视频到本地
		_, err = io.Copy(outplay, file)
		if err != nil {
			return
		}

		// 对视频进行压缩
		cmdplay := exec.Command("ffmpeg", "-i", `C:\Users\84023\Videos\`+fornowplay, "-r", "30", "-b:v", "2M", "-b:a", "1.5M", `C:\Users\84023\Videos\`+play)
		err = cmdplay.Run()
		if err != nil {
			return
		}
		outplay.Close() // 因为后续需要删除文件，所以需要提前关闭

		// 删除未压缩的视频版本
		err = os.Remove(`C:\Users\84023\Videos\` + fornowplay)
		if err != nil {
			return
		}

		// 抽取视频的第一帧作为图片
		cmdpic := exec.Command("ffmpeg", "-i", `C:\Users\84023\Videos\`+play, "-y", "-f", "image2", "-vframes", "1", `C:\Users\84023\Pictures\images\`+cover)
		err = cmdpic.Run()
		if err != nil {
			return
		}
		file.Close()

		// 用户数据表的作品+1
		// 异步操作
		// err = db.Model(&userinfo).Where("ID = ?", userinfo.ID).Update("WorkCount", userinfo.WorkCount+1).Error
		// if err != nil {
		// 	fmt.Printf("user updating error")
		// }
		go dbcreate(db, userinfo)

		// 用户缓存作品+1
		_, err = redisDb.HIncrBy("user_"+id_str, "work_count", 1).Result()
		if err != nil {
			fmt.Println("redis increase error")
			return
		}

		l := video.NewPublishLogic(r.Context(), svcCtx)
		resp, err := l.Publish(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
