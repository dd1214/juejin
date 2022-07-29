package controller

/*
import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	newID := makeId()
	publishTime := time.Now().Unix()
	dbInit()
	defer db.Close()
	var users []dbUser
	//查询
	db.Select(&users, "select ID, Name from User where token=?", token)
	if users == nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//获取视频标题
	title := c.PostForm("title")
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", users[0].ID, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//视频截取封面图，保存后返回图片名
	pegName := finalName[:len(finalName)-4] + "cover"
	snapshotName := GetSnapshot("./public/"+finalName, "./public/"+pegName, 1)

	db.Exec("insert into Video(ID, AuthorID, PlayUrl, CoverUrl, FavoriteCount, CommentCount, IsFavorite, Title, PublishTime)value(?, ?, ?, ?, ?, ?, ?, ?, ?)", newID, users[0].ID, "http://192.168.123.32:8080/static/"+finalName, "http://192.168.123.32:8080/static/"+snapshotName, 0, 0, 0, title, publishTime)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userID := c.Query("user_id")

	dbInit()
	defer db.Close()
	var videoList []Video
	//获取用户信息
	var user []dbUser
	db.Select(&user, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", userID)
	if user == nil {
		return
	}
	var ResponseUser = User{
		Id:            user[0].ID,
		Name:          user[0].Name,
		FollowCount:   user[0].FollowCount,
		FollowerCount: user[0].FollowerCount,
		IsFollow:      user[0].IsFollow,
	}
	//获取视频列表
	var videos []dbVideo
	db.Select(&videos, "select ID, PlayUrl, CoverUrl, FavoriteCount, CommentCount, IsFavorite, Title from Video where AuthorID=?", user[0].ID)
	//填充视频列表
	for _, video := range videos {
		videoList = append(videoList, Video{
			Id:            video.ID,
			Author:        ResponseUser,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
		})
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		VideoList: videoList,
	})
}

// GetSnapshot 生成视频缩略图并保存（作为封面）
// 视频路径: videoPath, 生成的略缩图保存路径: snapshotPath, 略缩图所属帧数: frameNum, 返回值为略缩图文件名: snapshotName
func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	err = imaging.Save(img, snapshotPath+".jpeg")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	// 成功则返回生成的缩略图名
	names := strings.Split(snapshotPath, "/")
	snapshotName = names[len(names)-1] + ".jpeg"
	return
}
*/
