package cmd

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/tidwall/gjson"

	"github.com/spf13/cobra"
)

var (
	flagOpen bool
)

var translateCmd = &cobra.Command{
	Use:     "translate",
	Aliases: []string{"fanyi", "fy"},
	Run: func(cmd *cobra.Command, args []string) {
		text := args[0]

		if flagOpen {
			link := fmt.Sprintf("http://dict.youdao.com/search?q=%s&ue=utf8&keyfrom=cli", text)
			if err := openLink(link); err == nil {
				return
			}
		}

		translatedText, err := youdaoTanslate(text)
		if err != nil {
			fmt.Println("Don't know how to translate: ", text)
			return
		}

		fmt.Print(translatedText)
		return
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)
	translateCmd.Flags().BoolVarP(&flagOpen, "open", "o", false, "open link")
}

// youdao ranslate
func youdaoTanslate(text string) (string, error) {
	appid := "5efe212bb3d9956b"
	key := "t4vaYZe0ojVXzO6quun0Nc1WbMlE1wK9"
	nonce := strconv.FormatInt(time.Now().Unix(), 10)

	str := fmt.Sprintf("%s%s%s%s", appid, text, nonce, key)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(str)))

	req := url.Values{}
	req.Set("q", text)
	req.Set("from", "auto")
	req.Set("to", "auto")
	req.Set("appKey", appid)
	req.Set("salt", nonce)
	req.Set("sign", sign)

	reqURL := "https://openapi.youdao.com/api?" + req.Encode()
	res, err := http.Get(reqURL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var trans1, trans2 string

	data := gjson.ParseBytes(body)
	arr1 := data.Get("basic.explains").Array()
	if len(arr1) > 0 {
		trans1 = fmt.Sprintf("英汉翻译:\n")
		for _, v := range arr1 {
			trans1 += fmt.Sprintf("%s\n", v)
		}
	}

	arr2 := data.Get("web").Array()
	if len(arr2) > 0 {
		trans2 = fmt.Sprintf("\n网络释义:\n")
		for _, v := range arr2 {
			key := v.Get("key").String()
			value := v.Get("value").Array()
			trans2 += fmt.Sprintf("%s: %s\n", key, value[0])
		}
	}

	if trans1 == "" && trans2 == "" {
		return "", fmt.Errorf("Don't know how to translate: %s", text)
	}

	trans := trans1 + trans2
	return trans, nil
}
