package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"wechat/global"
	"wechat/model"
)

type DownLoadService struct {
}

func (ds *DownLoadService) GetDownLoadImages(page int) {
	var bookList []model.ChineseBook
	bookDB := global.GVA_DB.Model(&model.ChineseBook{}).Debug()
	bookDB.Find(&bookList)
	log.Printf("bookList count:%+v \n", len(bookList))
	var urlList []string
	for idx, _ := range bookList {
		urlList = append(urlList, bookList[idx].Icon)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(urlList))

	for idx, _ := range urlList {
		go ds.DoLoadImg(urlList[idx], &wg)
	}

	wg.Wait()
}

func (ds *DownLoadService) DoLoadImg(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	var outputPath string = "/Users/jiang/Documents/picture/chinese_book/cover"
	ds.DownloadFileSavePath(url, outputPath)
}

//DownloadFileSavePath 下载文件，如图片、MP3等
func (ds *DownLoadService)DownloadFileSavePath(url, directory string) error {
	// 获取文件名
	fileName := ds.ExtractFileName(url)

	// 构建完整的保存路径
	savePath := filepath.Join(directory, fileName)

	// 创建目录（如果不存在）
	err := os.MkdirAll(directory, 0755)
	if err != nil {
		return err
	}

	// 发起 HTTP 请求并下载文件
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	// 创建保存文件
	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将下载的内容保存到文件中
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

//ExtractFileName 截取最后一个"/"后面的内容
func (ds *DownLoadService) ExtractFileName(url string) string {
	// 使用 strings.LastIndex 找到最后一个斜杠的位置
	lastSlashIndex := strings.LastIndex(url, "/")
	if lastSlashIndex == -1 {
		return "" // 没有找到斜杠，返回空字符串或者可以进行错误处理
	}

	// 截取最后一个斜杠之后的部分作为文件名
	fileName := url[lastSlashIndex+1:]

	return fileName
}
