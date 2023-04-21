# epub翻译器

一个翻译epub书籍的工具

# 特征

-   [x] 翻译 epub 书籍
-   [x] 使用谷歌翻译 API
-   [ ] 使用 OpenAI GPT-3.5 API
-   [ ] 使用 DeepL API
-   [ ] 翻译文本文件
-   [ ] 翻译 PDF 文件

# 用法

## 安装

下载发布自[发布页面](https://github.com/smark-d/epub-translator/releases)基于您的平台。

## 跑步

> 它需要 sudo 权限才能运行。因为它需要创建和删除一个目录`./temp`目录。

```bash
chmod +x epub-translator
sudo ./epub-translator -f ./path/xxx.epub -s en -t zh-CN -e google
```

After the translation is completed, the translated epub file will be generated in the same directory named `xx.epub.translated`.

## 选项

```bash
Usage of ./epub-translator:
  -e string
        engine: google, openai, deepl (default "google")
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
