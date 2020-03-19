package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ini=`
;ini文件读写示例
[core]
repositoryformatversion = 0
filemode = false
bare = false
logallrefupdates = true
symlinks = false
ignorecase = true
hideDotFiles = dotGitOnly

[remote "origin"]
url = https://github.com/davyxu/cellnet
fetch = +refs/heads/*:refs/remotes/origin/*

[branch "master"]
remote = origin
merge = refs/heads/master
`

// 根据文件名，段名，键名获取ini的值
func getValue(filename, expectSection, expectKey string) string {
// 文件句柄
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	// 缓存读写器
	reader := bufio.NewReader(file)

	// 当前读取的段的名字
	var sectionName string
	for {
		// 读取一行
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// 切掉行的左右两边的空白字符
		linestr = strings.TrimSpace(linestr)
		// 忽略空行
		if linestr == "" {
			continue
		}

		// 忽略注释
		if linestr[0] == ';' {
			continue
		}

		// 行首和尾巴分别是方括号的，说明是段标记的起止符
		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
			// 将段名取出
			sectionName = linestr[1 : len(linestr)-1]
			// 这个段是希望读取的
		} else if sectionName == expectSection {
			// 切开等号分割的键值对
			pair := strings.Split(linestr, "=")
			// 保证切开只有1个等号分割的简直情况
			if len(pair) == 2 {
				// 去掉键的多余空白字符
				key := strings.TrimSpace(pair[0])
				// 是期望的键
				if key == expectKey {
					// 返回去掉空白字符的值
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(getValue("config.ini", "remote \"origin\"", "fetch"))
	fmt.Println(getValue("config.ini", "core", "hideDotFiles"))
}