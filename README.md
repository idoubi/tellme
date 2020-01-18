## tellme 获取信息的命令行工具

### 安装说明

- MacOS 安装

```shell
brew tap idoubi/tools
brew install tellme
```

- 其他系统安装

下载源码，运行`go build`，运行生成的二进制文件。

### 使用示例

- 百科查找 baike

命令：`baike`、`bk`、`wiki`、`wk`

```shell
tellme baike 周杰伦
# 百度百科查找

tellme baike 周杰伦 -p wp
# 维基百科查找

tellme baike 周杰伦 -p hd
# 互动百科查找
```

- 搜索 search

命令：`search`、`sr`

```shell
tellme search 周杰伦
# google搜索

tellme search 周杰伦 -p bd
# 百度搜索

tellme search 周杰伦 -p wx
# 微信搜索

tellme search 周杰伦 -p zh
# 知乎搜索

tellme search 周杰伦 -p jj
# 掘金搜索
```

- 翻译 translate

命令：`translate`、`fanyi`、`fy`

```shell
tellme fanyi 你好
# 中译英

tellme fanyi "hello world"
# 英译中

tellme fanyi -o "what's your name"
# 打开浏览器查看释义
```

- 日期时间 time

```shell
tellme time 
# 获取当前时间。输出：2020-01-10 14:23:10 己亥猪年腊月十六日 Friday

tellme time -s
# 获取当前时间串。输出：2020-01-10 13:49:46

tellme time -s 1548635411
# 将指定时间戳转换为时间串。输出：2019-01-28 08:30:11

tellme time -t 
# 获取当前时间戳。输出：1578635437

tellme time -t "2020-02-02 20:00:02"
# 将指定时间戳转换为时间戳。输出：1580673602

tellme time -w
# 查看今天是星期几。输出：Friday

tellme time -w "2020-02-02"
# 查看指定时间是星期几。输出：Sunday

tellme time -n
# 查看今天的农历日期。输出：己亥猪年腊月十六日

tellme time -n "2020-02-02"
# 查看指定时间的农历日期。输出：庚子鼠年正月初九日

tellme time -y "2020-01-01"
# 指定农历日期查看阳历日期。输出：2020-01-25

tellme time -y -r "1993-03-17"
# 指定闰年农历日期查看阳历日期。输出：1993-05-08

tellme time -j
# 查看今天是什么节日。输出：[中国110宣传日]

tellme time -j "2020-04-01"
# 查看指定日期是什么节日。输出：[愚人节]

tellme time -o
# 原样输出当前时间。输出：2020-01-10 14:22:27.890235 +0800 CST m=+0.000376593

tellme time -o "2020-02-02 15:22:11"
# 原样输出指定时间。输出：2020-02-02 15:22:11 +0000 UTC

tellme time -f "2006/01/02 15:04:05"
# 按指定格式输出当前时间。输出：2020/01/10 13:51:54

tellme time -f "2006/01/02" 1578615616
# 按指定格式输出指定时间。输出：2020/01/10
```