package filesha1

import (
	"crypto/sha1"
	"encoding/json"
	"filesha1/pkg/log"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

type FileSha1 struct {
	Root           string   `json:"root"`
	OutputFileName string   `json:"outputFile"`
	Exclude        []string `json:"exclude"`
	OutFile        *os.File
}

func NewFileSha1(config string) (fileSha1 *FileSha1, err error) {
	fileSha1 = &FileSha1{}
	log.Debug("配置%s", config)
	err = json.Unmarshal([]byte(config), fileSha1)
	return fileSha1, err
}

/*
遍历目录
*/
func (c *FileSha1) HandleFilelist() error {
	c.initialize()
	err := filepath.Walk(c.Root, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if c.isExclude(path) {
			return nil
		}
		c.GenerateSha1(path, f)
		return nil
	})
	if c.OutFile != nil {
		c.OutFile.Close()
	}
	return err
}
func (c *FileSha1) initialize() {
	//绝对路径
	//	currPath, error1 := filepath.Abs(c.Root)
	//	if error1 != nil {
	//		log.Error(error1.Error())
	//	}
	// 清理路径中的多余字符,并给匹配符的前面加上路径分割符
	for index, _ := range c.Exclude {
		//转换/为对应系统的分割符
		filepath.FromSlash(c.Exclude[index])
		filepath.Clean(c.Exclude[index])
		//		c.Exclude[index] = c.Root + string(os.PathSeparator) + "*" + string(os.PathSeparator) + c.Exclude[index] + "*"
		c.Exclude[index] = string(os.PathSeparator) + c.Exclude[index]
	}
	// 处理输出文件
	var err1 error
	filepath.FromSlash(c.OutputFileName)
	filepath.Clean(c.OutputFileName)
	if checkFileIsExist(c.OutputFileName) { //如果文件存在
		c.OutFile, err1 = os.OpenFile(c.OutputFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend) //打开文件
		if err1 != nil {
			log.Error(err1.Error())
		}
	} else {
		c.OutFile, err1 = os.Create(c.OutputFileName) //创建文件
		if err1 != nil {
			log.Error(err1.Error())
		}
	}
}

/*
*判断是否过滤
 */
func (c *FileSha1) isExclude(path string) bool {
	for _, value := range c.Exclude {
		//		ok, err := filepath.Match(value, path)
		ok, err := regexp.Match(value, []byte(path))
		if err != nil {
			log.Error(err.Error())
			continue
		}
		if ok {
			log.Debug("过滤路径： %s, 过滤规则： %s", path, value)
			return true
		}
	}
	return false
}

/**
生成sha1
*/
func (c *FileSha1) GenerateSha1(path string, f os.FileInfo) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	h := sha1.New()
	_, erro := io.Copy(h, file)
	if erro != nil {
		return
	}
	sha1StringInfo := fmt.Sprintf("path: %s, sha1: %x, size: %d \n", path, h.Sum(nil), f.Size())
	_, err = io.WriteString(c.OutFile, sha1StringInfo)
	if err != nil {
		log.Error(err.Error())
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
