package chapter2

import (
	"fmt"
	"testing"

	"github.com/dip-dev/go-test-tutorial/chapters/chapter2/communication"

	"github.com/golang/mock/gomock"
)

func TestExecWithMock(t *testing.T) {
	success := map[string]struct {
		want string
	}{
		"mock化テスト": {
			want: "Nice to meet you!!",
		},
	}

	for tt, tc := range success {
		t.Run(tt, func(t *testing.T) {
			fmt.Println(tc.want) // このfmtは一時的なエラー対応のため、テスト実装後に削除してください。

			// FIXME: mock contoroller生成
			ctrl := gomock.NewController(t)
			// FIXME: mock設定
			defer ctrl.Finish()
			// FIXME: mockを構造体(Chapter2)に設定
			mock := communication.NewMockInterfaceCommunication(ctrl)
			// FIXME: exec関数を呼び出し
			mock.EXPECT().Greeting().Return("Hello!")
			// FIXME: 結果を検証
			sample := New(mock)
			sample.exec()
		})
	}
}
