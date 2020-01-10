package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	flagBaikePlatform string
)

var openCmds = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

var baikeCmd = &cobra.Command{
	Use:     "baike",
	Aliases: []string{"bk", "wk", "wiki"},
	Short:   "find things in baike site",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := findInBaike(args[0], flagBaikePlatform)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(baikeCmd)
	baikeCmd.Flags().StringVarP(&flagBaikePlatform, "platform", "p", "baidu", "platform to find things")
}

// 百科查找
func findInBaike(keyword, platform string) error {
	var link string
	// 百度百科搜索
	if platform == "baidu" || platform == "bd" {
		link = fmt.Sprintf("https://baike.baidu.com/item/%s", keyword)
	}
	// 互动百科搜索
	if platform == "hudong" || platform == "baike" || platform == "hd" {
		link = fmt.Sprintf("http://www.baike.com/wiki/%s", keyword)
	}
	// 维基百科搜索
	if platform == "wikipedia" || platform == "wiki" || platform == "wp" {
		link = fmt.Sprintf("https://zh.wikipedia.org/zh-hans/%s", keyword)
	}
	if link == "" {
		return fmt.Errorf("invalid platform")
	}
	goos := runtime.GOOS
	opencmd := "open"
	opencmd, ok := openCmds[goos]
	if !ok {
		return fmt.Errorf("can not open link in %s", goos)
	}
	if err := exec.Command(opencmd, link).Start(); err != nil {
		return err
	}

	return nil
}
