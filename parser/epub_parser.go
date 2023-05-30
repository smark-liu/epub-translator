package parser

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/smark-d/epub-translator/trans"
)

type EpubParser struct {
	Path       string // epub file Path
	From       string // source language
	To         string // target language
	KeepOrigin bool   // keep the original file
	Translator trans.Translator
	tempDir    string
}

func (e *EpubParser) Parse() (string, error) {
	// Unzip the epub file.
	e.tempDir = filepath.Join("./temp", "parser", time.Now().Format("20060102150405"))
	err := e.unzip(e.Path, e.tempDir)
	if err != nil {
		log.Fatalf("Unzip the epub file %s error: %v", e.Path, err)
		panic(err)
	}
	log.Printf("Unzip the epub file %s to %s\n", e.Path, e.tempDir)
	e.translate()
	e.zip()
	return e.Path + ".chinese", nil
}

func (e *EpubParser) translate() {
	wg := &sync.WaitGroup{}
	filepath.Walk(e.tempDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && (filepath.Ext(path) == ".html" || filepath.Ext(path) == ".xhtml" || filepath.Ext(path) == ".htm") {
			// Open the file and read its contents.
			file, err := os.OpenFile(path, os.O_RDWR, 0666)
			if err != nil {
				panic(err)
			}
			wg.Add(1)
			go e.translateFile(file, wg)
		}
		return nil
	})
	wg.Wait()
}

func (e *EpubParser) translateFile(file *os.File, wg *sync.WaitGroup) {
	defer file.Close()
	defer wg.Done()
	reader, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		panic(err)
	}
	// find all <p> tags
	nodeP := reader.Find("p")
	for i := 0; i < nodeP.Size(); i++ {
		// translate the text and insert it after the <p> tag
		translate, _ := e.Translator.Translate(nodeP.Eq(i).Text(), e.From, e.To)
		nodeP.Eq(i).AfterHtml("<p>" + translate + "</p>")
		log.Printf("Translate the file: %s text: %s\n", file.Name(), nodeP.Eq(i).Text())
	}

	// save and replace the source file
	ret, err := reader.Html()
	if err != nil {
		panic(err)
	}
	// set the file pointer to the beginning of the file
	_, err = file.Seek(0, 0)
	n, err := file.WriteString(ret)
	log.Printf("Translate the chepter %s, wrote %d bytes\n", file.Name(), n)
	if err != nil {
		panic(err)
	}

}

// zip zips the temporary directory to a new epub file.
func (e *EpubParser) zip() {
	zipFile, err := os.Create(e.Path + ".translated")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	filepath.Walk(e.tempDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// Get the path of the file relative to the temp directory
			// and create a new zip file.
			filePath, err := filepath.Rel(e.tempDir, path)
			if err != nil {
				panic(err)
			}
			zipFile, err := zipWriter.Create(filePath)
			if err != nil {
				panic(err)
			}

			// Open the file and copy its contents to the zip file.
			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			_, err = io.Copy(zipFile, file)
			if err != nil {
				panic(err)
			}
		}
		return nil
	})
}

func (e *EpubParser) unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
