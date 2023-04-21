# epub翻譯器

一個翻譯epub書籍的工具

# 特徵

-   [x] 翻譯 epub 書籍
-   [x] 使用谷歌翻譯 API
-   [ ] 使用 OpenAI GPT-3.5 API
-   [ ] 使用 DeepL API
-   [ ] 翻譯文本文件
-   [ ] 翻譯 PDF 文件

# 用法

## 安裝

下載發布自[發布頁面](https://github.com/smark-d/epub-translator/releases)基於您的平台。

## 跑步

> 它需要 sudo 權限才能運行。因為它需要創建和刪除一個目錄`./temp`目錄。

```bash
chmod +x epub-translator
sudo ./epub-translator -f ./path/xxx.epub -s en -t zh-CN -e google
```

翻譯完成後，會在同名目錄下生成翻譯後的epub文件`xx.epub.translated`.

## 選項

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

## Using proxy

> 如果您在中國，您可以使用代理訪問 Google Translate API。
> epub-translator 將使用`http_proxy`訪問 Google Translate API 的環境變量。

```bash
sudo env http_proxy=http://localhost:1087 ./epub-translator -f ./path/xxx.epub -s en -t zh-CN -e google
```

# 貢獻

-   克隆這個倉庫
-   跑步`go build`構建二進製文件
-   跑步`go test`測試代碼

歡迎公關。

# 執照

和
