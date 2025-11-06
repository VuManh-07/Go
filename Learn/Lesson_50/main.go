// String Functions
package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}

/*
Hàm,Cú pháp,Chức năng,Ví dụ
strings.Contains,"(s, substr string) bool",Kiểm tra xem s có chứa chuỗi con substr hay không.,"Contains(""Golang"", ""go"") → true"
strings.HasPrefix,"(s, prefix string) bool",Kiểm tra xem s có bắt đầu bằng chuỗi prefix hay không.,"HasPrefix(""Golang"", ""Go"") → false"
strings.HasSuffix,"(s, suffix string) bool",Kiểm tra xem s có kết thúc bằng chuỗi suffix hay không.,"HasSuffix(""file.txt"", "".txt"") → true"
strings.Index,"(s, substr string) int",Trả về vị trí index đầu tiên của substr trong s. Trả về -1 nếu không tìm thấy.,"Index(""banana"", ""an"") → 1"
strings.LastIndex,"(s, substr string) int",Trả về vị trí index cuối cùng của substr trong s.,"LastIndex(""banana"", ""an"") → 3"
strings.Replace,"(s, old, new string, n int) string",Thay thế n lần chuỗi old bằng chuỗi new trong s. Nếu n = -1 thì thay thế tất cả.,"Replace(""ohoho"", ""o"", ""a"", 2) → ""ahaho"""
strings.ReplaceAll,"(s, old, new string) string",Thay thế tất cả các lần xuất hiện của old bằng new.,"ReplaceAll(""ohoho"", ""o"", ""a"") → ""ahaha"""
strings.ToLower,(s string) string,Chuyển đổi tất cả các ký tự thành chữ thường.,"ToLower(""HELLO"") → ""hello"""
strings.ToUpper,(s string) string,Chuyển đổi tất cả các ký tự thành chữ hoa.,"ToUpper(""hello"") → ""HELLO"""
strings.TrimSpace,(s string) string,Xóa khoảng trắng (white space) ở đầu và cuối chuỗi.,"TrimSpace("" hello world "") → ""hello world"""
strings.Trim,"(s, cutset string) string",Xóa bất kỳ ký tự nào trong cutset ở đầu và cuối chuỗi s.,"Trim(""!?!hello?!?"", ""?!"") → ""hello"""
strings.Repeat,"(s string, count int) string",Trả về một chuỗi mới bao gồm s lặp lại count lần.,"Repeat(""ha"", 3) → ""hahaha"""
*/
