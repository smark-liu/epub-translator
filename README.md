# epub-translator
A tool for translating epub books

# Features

- [x] Translate epub books
- [x] Use Google Translate API
- [ ] Use OpenAI GPT-3.5 API
- [ ] Use DeepL API
- [ ] Translate text files
- [ ] Translate PDF files

# Usage

## Install

Download release from [release page](!https://github.com/smark-d/epub-translator/releases) based on your platform.

## Run

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
        engine: google, openai, deepl (default "google")
  -f string
        file filePath
  -s string
        sourceLanguage language
  -t string
        targetLanguage language
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