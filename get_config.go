package goutils

import (
	"os"
	"strings"
)

type Conf map[string][]string

func (conf *Conf) GetVal(key string) string {
	if vals, ok := (*conf)[key]; ok {
		return vals[0]
	}
	return ""
}

// read configure file
func GetConf(filename string) (Conf, error) {
	content, err := os.ReadFile(filename)
	if nil != err {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	conf := make(map[string][]string)
	// 遍历每一行
	for _, line := range lines {
		// 处理掉行首尾空字符
		line = strings.Trim(line, " \r\t")
		// “#”开头为注释
		if 0 == len(line) || '#' == line[0] {
			continue
		}
		// 键值对以“=”分割
		kv := strings.SplitN(line, "=", 2)
		// 处理掉空字符
		key := strings.Trim(kv[0], " \t")
		// 该key是否已存在
		vals, ok := conf[key]
		if !ok {
			vals = make([]string, 0)
			// 只有key没有值的情况创建空数组
			if len(kv) < 2 {
				conf[key] = vals
				continue
			}
		}
		// 处理掉空字符
		val := strings.Trim(kv[1], " \t")
		conf[key] = append(vals, val)
	}
	return conf, nil
}
