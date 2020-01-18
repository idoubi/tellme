package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	flagSearchPlatform string
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"sr"},
	Run: func(cmd *cobra.Command, args []string) {
		err := search(args[0], flagSearchPlatform)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&flagSearchPlatform, "platform", "p", "google", "search platform")
}

// 搜索
func search(keyword, platform string) error {
	var link string
	// 谷歌搜索
	if platform == "google" || platform == "gg" {
		link = fmt.Sprintf("https://www.google.com/search?q=%s", keyword)
	}
	// 百度搜索
	if platform == "baidu" || platform == "bd" {
		link = fmt.Sprintf("https://www.baidu.com/s?wd=%s", keyword)
	}
	// 微信搜索
	if platform == "wechat" || platform == "wx" {
		link = fmt.Sprintf("https://weixin.sogou.com/weixin?type=2&query=%s", keyword)
	}
	// 知乎搜索
	if platform == "zhihu" || platform == "zh" {
		link = fmt.Sprintf("https://www.zhihu.com/search?type=content&q=%s", keyword)
	}
	// 掘金搜索
	if platform == "juejin" || platform == "jj" {
		link = fmt.Sprintf("https://juejin.im/search?query=%s&type=all", keyword)
	}
	if link == "" {
		return fmt.Errorf("invalid platform")
	}

	return openLink(link)
}
