package parser

import (
	"github.com/smark-d/epub-translator/common"
	"testing"

	"github.com/smark-d/epub-translator/trans"
)

func TestEpubParser_Parse(t *testing.T) {
	type fields struct {
		Path       string
		From       string
		To         string
		Translator trans.Translator
		tempDir    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Test with epub file",
			fields: fields{
				Path:       "/home/smark/code/go/epub-translator/epub.epub",
				From:       common.EN,
				To:         common.ZH,
				Translator: trans.NewGoogleTranslator(),
				tempDir:    "",
			},
			want:    "test",
			wantErr: false,
		}, // Add this test case to the tests slice
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EpubParser{
				Path:       tt.fields.Path,
				From:       tt.fields.From,
				To:         tt.fields.To,
				Translator: tt.fields.Translator,
				tempDir:    tt.fields.tempDir,
			}
			got, err := e.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
