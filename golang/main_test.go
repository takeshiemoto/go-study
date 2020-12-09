package main // テストファイルはテスト関数と同じパッケージに置く

import "testing"

func TestDecode(t *testing.T) {
	post, err := decode("post.json") // テストされる関数を呼び出す
	if err != nil {
		t.Error(err)
	}
	if post.ID != 1 { // 結果が予想通りかチェック
		t.Error("Wrong id, was expecting 1 but got", post.ID)
	}
	if post.Content != "Hello world" {
		t.Error("Wrong content, was expecting 'Hello world' bud got", post.Content)
	}
}

func TestingEncode(t *testing.T) {
	t.Skip("skipping encoding for now") // テストをすべて省略
}
