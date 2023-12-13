package utils

import (
	"bytes"
	"compress/zlib"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"assemble/logger"

	"github.com/gin-gonic/gin"
)

var iw, _ = NewIdWorker(1)

func GoId() int64 {
	id, _ := iw.NextId()
	return id
}

// NowDateTime 系统时间
func NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// NowDate 系统日期
func NowDate() string {
	return time.Now().Format("2006-01-02")
}

// NowParseDate 字符串转日期
func NowParseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

// NowUnix 时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// Base64Encode 编码
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode 解码
func Base64Decode(str string) []byte {
	decoded, _ := base64.StdEncoding.DecodeString(str)
	return decoded
}

// IsString 判断字符串是否为空
func IsString(str string) bool {
	return strings.Trim(str, " ") == ""
}

func IsNotString(str string) bool {
	return !IsString(str)
}

// FixUrl 补全url
func FixUrl(mainSite *url.URL, nextLoc string) *url.URL {
	nextLocUrl, err := url.Parse(nextLoc)
	if err != nil {
		return nil
	}
	return mainSite.ResolveReference(nextLocUrl)
}

// Md5
// @param str 字符串
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// SHA256 算法
func SHA256(message []byte) string {
	bytes2 := sha256.Sum256(message)
	hashcode2 := hex.EncodeToString(bytes2[:])
	return hashcode2
}

// CreateFile 创建新文件
func CreateFile(filepath string) {
	//  如果没有文件目录就创建一个
	if _, err := os.Stat(filepath); err != nil {
		if !os.IsExist(err) {
			_ = os.MkdirAll(filepath, os.ModePerm)
		}
	}
}

// ZlibUnCompress 解压zlib
func ZlibUnCompress(compressSrc []byte) ([]byte, error) {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	_, err := io.Copy(&out, r)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// RandomNum 随机数
func RandomNum(num int32, max int32) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Int31n(max)
	if random > num {
		return int(random)
	}
	return RandomNum(num, max)
}

func SaveFile(fileDir string, file *multipart.FileHeader) (map[string]interface{}, error) {
	open, err := file.Open()
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(open)
	if err != nil {
		return nil, err
	}
	return SaveBufFile(fileDir, file.Filename, buf)
}

func SaveBufFile(fileDir, filename string, buf []byte) (map[string]interface{}, error) {
	temp, err := os.MkdirTemp(fileDir, "temp")
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"name": filename, "id": temp[strings.LastIndex(temp, "/")+1:]}, os.WriteFile(temp+"/"+filename, buf, 0666)
}

func FindFile(filename string, suffix map[string]struct{}) (string, error) {
	filename = GetSaveTmpFilePath() + filename
	if _, err := os.Stat(filename); err != nil {
		logger.Error("find dri fail ", err)
		return "", errors.New("no such file or directory")
	}

	var file string
	err := filepath.Walk(filename, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file = path
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if _, ok := suffix[filepath.Ext(file)]; ok {
		return file, nil
	}
	return "", errors.New(filename + "该文件后缀不符合要求")
}

func GetSaveTmpFilePath() string {
	return "/tmp/" + NowDate() + "/"
}

func SlicePage(page, pageSize, nums int) (sliceStart, sliceEnd int) {
	if pageSize > nums {
		return 0, nums
	}
	pageCount := int(math.Ceil(float64(nums) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}
	sliceStart = (page - 1) * pageSize
	sliceEnd = sliceStart + pageSize
	if sliceEnd > nums {
		sliceEnd = nums
	}
	return sliceStart, sliceEnd
}

// Slice2Set 切片转集合
func Slice2Set[T string | int64 | int](arr []T) map[T]struct{} {
	res := make(map[T]struct{})
	for _, item := range arr {
		res[item] = struct{}{}
	}
	return res
}

// DeduplicationStable 稳定去重
func DeduplicationStable[T string | int | int64](req []T) []T {
	n := len(req)
	if n == 0 {
		return req
	}

	set := make(map[T]struct{}, n)
	stable := make([]T, 0, n)

	for _, item := range req {
		if _, found := set[item]; !found {
			stable = append(stable, item)
		}
		set[item] = struct{}{}
	}

	return stable
}

func GetIP(c *gin.Context) string {
	// 获得当前登陆的ip
	nowIp := c.Request.Host
	if c.Request.Header.Get("X-Real-Ip") != "" {
		nowIp = c.Request.Header.Get("X-Real-Ip")
	}
	return strings.Split(nowIp, ":")[0]
}

func ArrayDates(start, end time.Time) []time.Time {
	var dates []time.Time
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	for start.Before(end) || start.Equal(end) {
		dates = append(dates, start)
		start = start.AddDate(0, 0, 1)
	}
	return dates
}
