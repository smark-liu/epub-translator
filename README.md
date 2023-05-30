# epub-translator
A tool for translating epub books [简体中文](./README.zh-CN.md)|[繁体中文](./README.zh-TW.md)

Examples:
![](./doc/img/translate.png)
# Features

- [x] Translate epub books
- [x] Use Google Translate API
- [x] Use OpenAI GPT-3.5 API
- [ ] Use DeepL API
- [ ] Translate text files
- [ ] Translate PDF files

# Usage

## Install

Download release from [release page](https://github.com/smark-d/epub-translator/releases) based on your platform.

## Config OpenAI API (Optional)
> If you don't want to use OpenAI API, you can skip this step.

1. Create an account on [OpenAI](https://openai.com/).
2. Create an API key on [OpenAI Dashboard](https://platform.openai.com/account/api-keys).
3. Create a file named `config.json` in the same directory as the binary file.
4. Copy the [config.example.json](./config.example.json) file and paste it into the config.json file.
5. Replace the `apiKey` with your API key.
6. Replace the apiUrl with your proxy url. (Optional, if you are in China, you can use this or http_proxy environment variable. see [Using proxy](#using-proxy)

## Run it

> It needs sudo permission to run. Because it needs to create and remove a directory in `./temp` directory.

```bash
chmod +x epub-translator
sudo ./epub-translator -f ./path/xxx.epub -s en -t zh-CN -e google
```

After the translation is completed, the translated epub file will be generated in the same directory named `xx.epub.translated`.

## Options

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
  -k bool
        keep the original text (default true)
```

## Using proxy

> If you are in China, you can use proxy to access Google Translate API.
> epub-translator will use `http_proxy` environment variables to access Google Translate API.

```bash
sudo env http_proxy=http://localhost:1087 ./epub-translator -f ./path/xxx.epub -s en -t zh-CN -e google
```

# Contributing

- clone this repo
- run `go build` to build the binary
- run `go test` to test the code

PR is welcome.

# License

MIT