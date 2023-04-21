# epub翻譯器

一個翻譯epub書籍的工具[簡體中文](./README.zh-CN.md)\|[繁體中文](./README.zh-TW.md)

# 特徵

-   [x] 翻譯 epub 書籍
-   [x] 使用谷歌翻譯 API
-   [x] 使用 OpenAI GPT-3.5 API
-   [ ] 使用 DeepL API
-   [ ] 翻譯文本文件
-   [ ] 翻譯 PDF 文件

# 用法

## 安裝

下載發布自[發布頁面](https://github.com/smark-d/epub-translator/releases)基於您的平台。

## 配置 OpenAI API（可選）

> 如果你不想使用 OpenAI API，你可以跳過這一步。

1.  創建一個帳戶[開放人工智能](https://openai.com/).
2.  創建 API 密鑰[OpenAI 儀表板](https://platform.openai.com/account/api-keys).
3.  創建一個名為`config.json`在與二進製文件相同的目錄中。
4.  複製[配置.example.json](./config.example.json)文件並將其粘貼到 config.json 文件中。
5.  更換`apiKey`使用您的 API 密鑰。
6.  將 apiUrl 替換為您的代理 url。 （可選，如果你在中國，可以使用這個或者http_proxy環境變量。見[使用代理](#using-proxy)

## 運行

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
        engine: google, openai (default "google")
  -f string
        file filePath
  -s string
        sourceLanguage language
  -t string
        targetLanguage language
```

## 使用代理

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
