package tools

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"time"
)

// NameOfFunction 获取函数的名字
func NameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// MD5 生成md5字符串
func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// GenerateRandCode 生成随机码 candidate 候选字符串 length 随机码长度
func GenerateRandCode(candidate string, length int) string {
	var letters = []rune(candidate)
	b := make([]rune, length)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// PathExists 判断所给路径文件/文件夹是否存在
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// ShellExec 执行本地 shell 命令
func ShellExec(cmd string) string {
	res := exec.Command("bash", "-c", cmd)
	output, err := res.CombinedOutput()
	if err != nil {
		return err.Error()
	}
	return string(output)
}
