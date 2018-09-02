/*
 * 说明：
 * 作者：zhe
 * 时间：2018-09-01 6:20 PM
 * 更新：
 */

package json_std

import "testing"

var data []byte

func BenchmarkNode_String(b *testing.B) {
	n := Node{
		Name:  "root",
		Value: 1,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data = []byte(n.String())
	}
}

func BenchmarkNode_Initialize(b *testing.B) {
	n := Node{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = n.Initialize(data)
	}
}
