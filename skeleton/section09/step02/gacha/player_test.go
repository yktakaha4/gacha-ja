// gachaパッケージとは別でgacha_testパッケージにする
package gacha_test

import (
	"testing"

	"github.com/gohandson/gacha-ja/gacha"
)

func TestPlayer_DrawableNum(t *testing.T) {
	cases := map[string]struct {
		tickets int
		coin    int
		want    int
	}{
		"zero-zero":      {0, 0, 0},
		"plus-zero":      {10, 0, 10},
		"plus-plus":      {10, 10, 11},
		"zero-plus":      {0, 10, 1},
		// コインが1回分に満たない場合のテスト
		"zero-short":     {0, 5, 0},
	}

	for name, tt := range cases {
		// ttをこのスコープで再定義しておく
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			p := gacha.NewPlayer(tt.tickets, tt.coin)
			got := p.DrawableNum()
			if got != tt.want {
				// 分かりやすいメッセージを出してテストを失敗させる
				t.Errorf("want is %v, got is %v", got, tt.want)
			}
		})
	}
}
