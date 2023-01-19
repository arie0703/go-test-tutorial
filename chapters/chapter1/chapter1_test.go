package chapter1

import (
	"testing"
)

func TestAddition(t *testing.T) {
	// 正常系のテストパターン
	success := map[string]struct {
		numA int
		numB int
		want int
	}{
		// FIXME: テストケースを追加
		"numAが最小値である": {
			numA: 0,
			numB: 30,
			want: 30,
		},
		"numBが最小値である": {
			numA: 30,
			numB: 10,
			want: 40,
		},
		"numAが最大値である": {
			numA: 100,
			numB: 30,
			want: 130,
		},
		"numBが最大値である": {
			numA: 30,
			numB: 200,
			want: 230,
		},
		"numA, numBが最小値である": {
			numA: 0,
			numB: 10,
			want: 10,
		},
		"numA, numBが最大値である": {
			numA: 100,
			numB: 200,
			want: 300,
		},
	}
	// エラー系のテストパターン
	fail := map[string]struct {
		numA       int
		numB       int
		wantErrStr string
	}{
		// FIXME: テストケースを追加
		"numAが最小値未満である場合": {
			numA:       -1,
			numB:       30,
			wantErrStr: "numAは0以上の数値を指定してください。",
		},
		"numBが最小値未満である場合": {
			numA:       30,
			numB:       9,
			wantErrStr: "numBは10以上の数値を指定してください。",
		},
		"numAが最大値を超す場合": {
			numA:       101,
			numB:       30,
			wantErrStr: "numAは100以下の数値を指定してください。",
		},
		"numBが最大値を超す場合": {
			numA:       30,
			numB:       201,
			wantErrStr: "numBは200以下の数値を指定してください。",
		},
		"numA, numBが最小値未満である場合": {
			numA:       -1,
			numB:       9,
			wantErrStr: "numAは0以上の数値を指定してください。",
		},
		"numA, numBが最大値を超す場合": {
			numA:       101,
			numB:       201,
			wantErrStr: "numAは100以下の数値を指定してください。",
		},
	}

	for tt, tc := range success {
		t.Run(tt, func(t *testing.T) {
			got, err := addtion(tc.numA, tc.numB)
			if err != nil {
				t.Errorf("err is not nil: %s", err)
			}
			if tc.want != got {
				t.Errorf("unexpected return. want:%d actual:%d", tc.want, got)
			}
		})
	}
	for tt, tc := range fail {
		t.Run(tt, func(t *testing.T) {
			got, err := addtion(tc.numA, tc.numB)
			if got != 0 {
				t.Errorf("unexpected return. want:0 actual:%d", got)
			}
			if tc.wantErrStr != err.Error() {
				t.Errorf("unexpected err. want:%s actual:%s", tc.wantErrStr, err)
			}
		})
	}
}
