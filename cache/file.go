package cache

import (
	"gitee.com/quant1x/gox/api"
	"os"
)

const (
	// 目录权限
	cacheDirMode os.FileMode = 0755
	// 文件权限
	cacheFileMode os.FileMode = 0644
	// 文件替换模式, 会用到os.TRUNC
	cacheReplace = os.O_CREATE | os.O_RDWR | os.O_TRUNC
	// 更新
	cacheUpdate = os.O_CREATE | os.O_WRONLY
)

// Touch 创建一个空文件
func Touch(filename string) error {
	_ = api.CheckFilepath(filename, true)
	return os.WriteFile(filename, nil, cacheFileMode)
}
