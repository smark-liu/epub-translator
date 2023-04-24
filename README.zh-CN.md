# epub翻译器

一个翻译epub书籍的工具[简体中文](./README.zh-CN.md)\|[繁体中文](./README.zh-TW.md)

例子：![](./doc/img/translate.png)

# 特征

-   [x] 翻译 epub 书籍
-   [x] 使用谷歌翻译 API
-   [x] 使用 OpenAI GPT-3.5 API
-   [ ] 使用 DeepL API
-   [ ] 翻译文本文件
-   [ ] 翻译 PDF 文件

# 用法

## 安装

下载发布自[发布页面](https://github.com/smark-d/epub-translator/releases)基于您的平台。

## 配置 OpenAI API（可选）

> 如果你不想使用 OpenAI API，你可以跳过这一步。

1.  创建一个帐户[开放人工智能](https://openai.com/).
2.  创建 API 密钥[OpenAI 仪表板](https://platform.openai.com/account/api-keys).
3.  创建一个名为`config.json`在与二进制文件相同的目录中。
4.  复制[配置.example.json](./config.example.json)文件并将其粘贴到 config.json 文件中。
5.  更换`apiKey`使用您的 API 密钥。
6.  将 apiUrl 替换为您的代理 url。 （可选，如果你在中国，可以使用这个或者http_proxy环境变量。见[使用代理](#using-proxy)

## 运行

> 它需要 sudo 权限才能运行。因为它需要创建和删除一个目录`./temp`目录。

```bash
chmod +x epub-translator
sudo ./epub-translator -f ./path/xxx.epub -s en -t zh-CN -e google
```

翻译完成后，会在同名目录下生成翻译后的epub文件`xx.epub.translated`.

## 选项

```bash
Usage of ./epub-translator:
  -e string
        engine: google, openai (default "google")
  -f string
        file filePath
  -s string
        sourceLanguage language
  -t string
        targetLanguage language
```

## 使用代理

> 如果您在中国，您可以使用代理访问 Google Translate API。
> epub-translator 将使用`http_proxy`访问 Google Translate API 的环境变量。

```bash
sudo env http_proxy=http://localhost:1087 ./epub-translator -f ./path/xxx.epub -s en -t zh-CN -e google
```

# 贡献

-   克隆这个仓库
-   跑步`go build`构建二进制文件
-   跑步`go test`测试代码

欢迎公关。

# 执照

和
