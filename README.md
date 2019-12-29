## tellme 获取信息的命令行工具

### 基本命令

```shell
命令行获取信息

Usage:
  tellme [flags]
  tellme [command]

Available Commands:
  baike       在百科中查找
  help        查看使用说明
  version     查看版本

Flags:
  -h, --help      查看使用说明
  -v, --version   查看版本

Use "tellme [command] --help" 查看子命令使用说明
```

### 百科查找命令

```shell
在百科中查找

Usage:
  tellme baike [flags]

Aliases:
  baike, bk, wk, wiki

Flags:
  -h, --help              查看使用说明
  -p, --platform string   指定百科平台 (默认为百度百科)
                          bd 百度百科
                          hd 互动百科
                          wp 维基百科
```