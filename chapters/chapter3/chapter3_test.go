package chapter3

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/dip-dev/go-test-tutorial/mysql"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/dip-dev/go-test-tutorial/mysql/queries"
	"github.com/dip-dev/go-test-tutorial/mysql/structs"
)

var mysqlCli *mysql.Client

// TestMain(m *testing.M)については README.md を参照してください
func TestMain(m *testing.M) {
	// 前処理 start

	cli, err := mysql.New()
	if err != nil {
		log.Fatalf("[FATAL] %+v", err)
	}
	mysqlCli = cli

	os.Exit(m.Run())
}

// FIXME: ↓↓テストを作成する↓↓
func TestSelectPrefecture(t *testing.T) {
	success := map[string]struct {
		name string
		want *structs.MPrefecture
	}{
		"検索成功の場合": {
			name: "東京都",
			want: &structs.MPrefecture{
				ID:       13,
				Name:     "東京都",
				Area:     "関東",
				LandArea: 2194,
			},
		},
		"nilの場合": {
			name: "hoge",
			want: nil,
		},
	}

	fail := map[string]struct {
		name string
		want string
	}{
		"クエリエラーの場合": {
			name: "Error",
			want: "Error 1146: Table 'tutorial_user.hogehoge' doesn't exist",
		},
	}

	t.Run("成功", func(t *testing.T) {
		chapter3 := New(mysqlCli, queries.New())
		for tt, tc := range success {
			t.Run(tt, func(t *testing.T) {

				got, err := chapter3.selectPrefecture(tc.name)
				fmt.Println(tc.name)
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			})
		}
	})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("失敗", func(t *testing.T) {
		for tt, tc := range fail {
			t.Run(tt, func(t *testing.T) {
				mock := queries.NewMockSelecters(ctrl)

				// 有効でないクエリを投げる
				mock.EXPECT().SelectPrefecture().Return("SELECT id, name, area, land_area FROM hogehoge")
				chapter3 := New(mysqlCli, mock)

				got, err := chapter3.selectPrefecture(tc.name)
				assert.Nil(t, got)
				assert.EqualError(t, err, tc.want)
			})
		}
	})
}
