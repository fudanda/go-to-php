package str

import (
	"testing"
)

// Pimplode使用案例
func TestPimplode (t *testing.T) {
	// 	切片
	test := []string{"编", "程"}
	// 数组
	// test := [3]string{"a", "b"}
	str := Pimplode(",", test)
	t.Log(str)
}

// PJoin使用案例，同Pimplode
func TestPJoin(t *testing.T) {
	// 	切片
	test := []string{"编", "程"}
	// 数组
	// test := [3]string{"a", "b"}
	str := Pjoin(",", test)
	t.Log(str)
}

// Pexplode使用案例
func TestPexplode(t *testing.T) {
	test := "hello,world"
	// limit = 0
	slice := Pexplode(",", test, 0)
	t.Log(slice)

	// limit < 0
	slice = Pexplode(",", test, -1)
	t.Log(slice)

	// limit >0
	slice = Pexplode(",", test, 6)
	t.Log(slice)
}

// Pstrlen使用案例
func TestPstrlen(t *testing.T) {
	str := "1q测试"
	byteNum := Pstrlen(str)
	t.Log(byteNum)
}

// Pmb_strlen使用案例
func TestPmb_strlen(t *testing.T) {
	str := "1q测试"
	charNum := Pmb_strlen(str)
	t.Log(charNum)
}

// Plcfirst使用案例
func TestPlcfirst(t *testing.T) {
	str := " "
	lowerStr := Plcfirst(str)
	t.Log(lowerStr)
}

// Pucfirst使用案例
func TestPucfirst(t *testing.T) {
	str := "a b c"
	upperStr := Pucfirst(str)
	t.Log(upperStr)
}

// Pstrtoupper使用案例
func TestPstrtoupper(t *testing.T) {
	str := "1aB编程"
	uStr := Pstrtoupper(str)
	t.Log(uStr)
}

// Pstrtolower使用案例
func TestPstrtolower(t *testing.T) {
	str := "1AS编程"
	uStr := Pstrtolower(str)
	t.Log(uStr)
}

// Pucword使用案例
func TestPucword(t *testing.T) {
	str := "a 1 B 编 and 程"
	uStr := Pucword(str)
	t.Log(uStr)
}

// Ptrim使用案例
func TestPtrim(t *testing.T) {
	str := "z2,"
	trimStr := Ptrim(str, ",")
	t.Log(trimStr)
}

// Pltrim使用案例
func TestPltrim(t *testing.T) {
	str := " ltrim()"
	character_mask := ""
	// character_mask = "r"
	ltrimStr := Pltrim(str, character_mask)
	t.Log(ltrimStr, len(str), len(ltrimStr))
}

// Prtrim使用案例
func TestPrtrim(t *testing.T) {
	str := "rtrim() "
	character_mask := ""
	// character_mask = "r"
	rtrimStr := Prtrim(str, character_mask)
	t.Log(rtrimStr, len(str), len(rtrimStr))
}

// Pchop使用案例
func TestPchop(t *testing.T) {
	str := "chop()是rtrim()的别名  "
	t.Log(Pchop(str, " "))
}

// Pmd5使用案例
func TestPmd5(t *testing.T) {
	str := ""
	t.Log(Pmd5(str))
}

// Psha1使用案例
func TestPsha1(t *testing.T) {
	str := ""
	t.Log(Psha1(str))
}

// Pord使用案例
func TestPord(t *testing.T) {
	str := " "
	t.Log(Pord(str))
}

// Pallord使用案例
func TestPallord(t *testing.T) {
	str := "   "
	t.Log(Pallord(str))
}

// Pchr使用案例
func TestPchr(t *testing.T) {
	t.Log(Pchr(9711))
}

// Pecho使用案例
func TestPecho(t *testing.T) {
	Pecho(1, "232", "cdc")
	Pecho("heellp")
}

// Pvar_dump()使用案例
func TestPvar_dump(t *testing.T) {
	Pvar_dump(1, "h")
}

// Pprint使用案例
func TestPprint(t *testing.T) {
	Pprint("hello world")
}

// Pstr_repeat使用案例
func TestPstr_repeat(t *testing.T) {
	t.Log(Pstr_repeat("a", -1))
}
