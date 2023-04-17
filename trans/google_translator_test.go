package trans

import (
	"context"
	"testing"
)

func TestGoogleTranslator_Translate(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		source     string
		sourceLang string
		targetLang string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Translation",
			fields: fields{
				ctx: context.Background(),
			},
			args: args{
				source:     "Hello",
				sourceLang: "en",
				targetLang: "zh-CN",
			},
			want:    "你好",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGoogleTranslator()
			got, err := g.Translate(tt.args.source, tt.args.sourceLang, tt.args.targetLang)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoogleTranslator.Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GoogleTranslator.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
