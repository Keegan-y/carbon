package carbon

import (
	"fmt"
	"testing"
	"time"
)

var TimezoneTests = []struct {
	input    string // 输入值
	timezone string // 输入参数
	output   string // 期望输出值
}{
	{"2020-08-05 13:14:15", PRC, "2020-08-05 13:14:15"},
	{"2020-08-05", Tokyo, "2020-08-05 01:00:00"},
	{"2020-08-05", "Hangzhou", "panic"}, // 异常情况
}

func TestCarbon_Timezone1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("catch an exception in Timezone()：%s\n", r)
		}
	}()

	for _, v := range TimezoneTests {
		output := Timezone(v.timezone).Parse(v.input).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_Timezone2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("catch an exception in Timezone()：%s\n", r)
		}
	}()

	for _, v := range TimezoneTests {
		output := Timezone(PRC).Timezone(v.timezone).Parse(v.input).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_Now(t *testing.T) {
	expected := time.Now().Format(DateTimeFormat)

	output := Now().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}

	output = Timezone(PRC).Now().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}

}

func TestCarbon_Yesterday(t *testing.T) {
	expected := time.Now().AddDate(0, 0, -1).Format(DateTimeFormat)

	output := Yesterday().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}

	output = Timezone(PRC).Yesterday().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}
}

func TestCarbon_Tomorrow(t *testing.T) {
	expected := time.Now().AddDate(0, 0, 1).Format(DateTimeFormat)

	output := Tomorrow().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}

	output = Timezone(PRC).Tomorrow().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}
}

func TestCarbon_BeginningOfYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-01 00:00:00"},
		{"2020-02-01 13:14:15", "2020-01-01 00:00:00"},
		{"2020-02-28", "2020-01-01 00:00:00"},
		{"2020-02-29", "2020-01-01 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).BeginningOfYear().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-12-31 23:59:59"},
		{"2020-01-31 23:59:59", "2020-12-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-12-31 23:59:59"},
		{"2020-02-28", "2020-12-31 23:59:59"},
		{"2020-02-29", "2020-12-31 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfYear().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_BeginningOfMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-01 00:00:00"},
		{"2020-02-01 13:14:15", "2020-02-01 00:00:00"},
		{"2020-02-28", "2020-02-01 00:00:00"},
		{"2020-02-29", "2020-02-01 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).BeginningOfMonth().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-31 23:59:59"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-29 23:59:59"},
		{"2020-02-28", "2020-02-29 23:59:59"},
		{"2020-02-29", "2020-02-29 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfMonth().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_BeginningOfWeek(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2019-12-30 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-27 00:00:00"},
		{"2020-02-01 13:14:15", "2020-01-27 00:00:00"},
		{"2020-02-28", "2020-02-24 00:00:00"},
		{"2020-02-29", "2020-02-24 00:00:00"},
		{"2020-10-04", "2020-09-28 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).BeginningOfWeek().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfWeek(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-05 23:59:59"},
		{"2020-01-31 23:59:59", "2020-02-02 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-02 23:59:59"},
		{"2020-02-28", "2020-03-01 23:59:59"},
		{"2020-02-29", "2020-03-01 23:59:59"},
		{"2020-10-04", "2020-10-04 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfWeek().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_BeginningOfDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-31 00:00:00"},
		{"2020-02-01 13:14:15", "2020-02-01 00:00:00"},
		{"2020-02-28", "2020-02-28 00:00:00"},
		{"2020-02-29", "2020-02-29 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).BeginningOfDay().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 23:59:59"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-01 23:59:59"},
		{"2020-02-28", "2020-02-28 23:59:59"},
		{"2020-02-29", "2020-02-29 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfDay().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_BeginningOfHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-31 23:00:00"},
		{"2020-02-01 13:14:15", "2020-02-01 13:00:00"},
		{"2020-02-28", "2020-02-28 00:00:00"},
		{"2020-02-29", "2020-02-29 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).BeginningOfHour().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:59:59"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-01 13:59:59"},
		{"2020-02-28", "2020-02-28 00:59:59"},
		{"2020-02-29", "2020-02-29 00:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfHour().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_BeginningOfMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:00"},
		{"2020-02-01 13:14:15", "2020-02-01 13:14:00"},
		{"2020-02-28", "2020-02-28 00:00:00"},
		{"2020-02-29", "2020-02-29 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).BeginningOfMinute().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:00:59"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-01 13:14:59"},
		{"2020-02-28", "2020-02-28 00:00:59"},
		{"2020-02-29", "2020-02-29 00:00:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfMinute().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_CreateFromTimestamp(t *testing.T) {
	Tests := []struct {
		timestamp int64  // 输入参数
		output    string // 期望输出值
	}{
		{1577855655, "2020-01-01 13:14:15"},
		{1580447655, "2020-01-31 13:14:15"},
		{1580534055, "2020-02-01 13:14:15"},
		{1582866855, "2020-02-28 13:14:15"},
		{1582953255, "2020-02-29 13:14:15"},
	}

	for _, v := range Tests {
		output := CreateFromTimestamp(v.timestamp).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).CreateFromTimestamp(v.timestamp).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_CreateFromDateTime(t *testing.T) {
	Tests := []struct {
		year, month, day, hour, minute, second int    // 输入参数
		output                                 string // 期望输出值
	}{
		{2020, 01, 01, 13, 14, 15, "2020-01-01 13:14:15"},
		{2020, 1, 31, 13, 14, 15, "2020-01-31 13:14:15"},
		{2020, 2, 1, 13, 14, 15, "2020-02-01 13:14:15"},
		{2020, 2, 28, 13, 14, 15, "2020-02-28 13:14:15"},
		{2020, 2, 29, 13, 14, 15, "2020-02-29 13:14:15"},
	}

	for _, v := range Tests {
		output := CreateFromDateTime(v.year, v.month, v.day, v.hour, v.minute, v.second).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).CreateFromDateTime(v.year, v.month, v.day, v.hour, v.minute, v.second).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_CreateFromDate(t *testing.T) {
	clock := Now().ToTimeString()

	Tests := []struct {
		year, month, day int    // 输入参数
		output           string // 期望输出值
	}{
		{2020, 01, 01, "2020-01-01 " + clock},
		{2020, 1, 31, "2020-01-31 " + clock},
		{2020, 2, 1, "2020-02-01 " + clock},
		{2020, 2, 28, "2020-02-28 " + clock},
		{2020, 2, 29, "2020-02-29 " + clock},
	}

	for _, v := range Tests {
		output := CreateFromDate(v.year, v.month, v.day).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).CreateFromDate(v.year, v.month, v.day).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_CreateFromTime(t *testing.T) {
	date := Now().ToDateString()

	Tests := []struct {
		hour, minute, second int    // 输入参数
		output               string // 期望输出值
	}{
		{0, 0, 0, date + " 00:00:00"},
		{00, 00, 15, date + " 00:00:15"},
		{00, 14, 15, date + " 00:14:15"},
		{13, 14, 15, date + " 13:14:15"},
	}

	for _, v := range Tests {
		output := CreateFromTime(v.hour, v.minute, v.second).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).CreateFromTime(v.hour, v.minute, v.second).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_CreateFromGoTime(t *testing.T) {
	Tests := []struct {
		time   time.Time // // 输入参数
		output string    // 期望输出值
	}{
		{time.Now(), time.Now().Format(DateTimeFormat)},
		{parseByLayout("2020-08-05 13:14:15", DateTimeFormat), "2020-08-05 13:14:15"},
	}

	for _, v := range Tests {
		output := CreateFromGoTime(v.time).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).CreateFromGoTime(v.time).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_Parse(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15"},
		{"20200805131415", "2020-08-05 13:14:15"},
		{"20200805", "2020-08-05 00:00:00"},
		{"2020-08-05", "2020-08-05 00:00:00"},
		{"2020-08-05T13:14:15+08:00", "2020-08-05 13:14:15"},
		{"12345678", "panic"}, // 异常情况
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("catch an exception in Parse()：%s\n", r)
		}
	}()

	for _, v := range Tests {
		output := Parse(v.input).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

var ParseByFormatTests = []struct {
	input  string // 输入值
	format string // 输入参数
	output string // 期望输出值
}{
	{"2020|08|05", "Y|m|d", "2020-08-05 00:00:00"},
	{"2020|08|05 13:14:15", "Y|m|d H:i:s", "2020-08-05 13:14:15"},
	{"12345678", "abc", "panic"}, // 异常情况
}

func TestCarbon_ParseByFormat1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("catch an exception in ParseByFormat()：%s\n", r)
		}
	}()

	for _, v := range ParseByFormatTests {
		output := ParseByFormat(v.input, v.format).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ParseByFormat2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("catch an exception in ParseByFormat()：%s\n", r)
		}
	}()

	for _, v := range ParseByFormatTests {
		output := Timezone(PRC).ParseByFormat(v.input, v.format).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

var ParseByDurationTests = []struct {
	input    string // 输入值
	duration string // 输入参数
	output   string // 期望输出值
}{
	{Now().ToDateTimeString(), "10h", time.Now().Add(parseByDuration("10h")).Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "-10h", time.Now().Add(parseByDuration("-10h")).Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "10.5h", time.Now().Add(parseByDuration("10.5h")).Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "-10.5h", time.Now().Add(parseByDuration("-10.5h")).Format(DateTimeFormat)},

	{Now().ToDateTimeString(), "10m", time.Now().Add(parseByDuration("10m")).Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "-10m", time.Now().Add(parseByDuration("-10m")).Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "10.5m", time.Now().Add(parseByDuration("10.5m")).Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "-10.5m", time.Now().Add(parseByDuration("-10.5m")).Format(DateTimeFormat)},

	{Now().ToDateTimeString(), "10s", time.Now().Add(parseByDuration("10s")).Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "-10s", time.Now().Add(parseByDuration("-10s")).Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "10.5s", time.Now().Add(parseByDuration("10.5s")).Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "-10.5s", time.Now().Add(parseByDuration("-10.5s")).Format(DateTimeFormat)},

	{Now().ToDateTimeString(), "-10a", "panic"}, // 异常情况
}

func TestCarbon_ParseByDuration1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("catch an exception in ParseByDuration()：%s\n", r)
		}
	}()

	for _, v := range ParseByDurationTests {
		output := ParseByDuration(v.duration).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_ParseByDuration2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("catch an exception in ParseByDuration()：%s\n", r)
		}
	}()

	for _, v := range ParseByDurationTests {
		output := Timezone(PRC).ParseByDuration(v.duration).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_Duration(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		duration string // 输入参数
		output   string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "10h", "2020-01-01 23:14:15"},
		{"2020-01-01 13:14:15", "-10h", "2020-01-01 03:14:15"},
		{"2020-01-01 13:14:15", "10.5h", "2020-01-01 23:44:15"},
		{"2020-01-01 13:14:15", "-10.5h", "2020-01-01 02:44:15"},

		{"2020-01-01 13:14:15", "10m", "2020-01-01 13:24:15"},
		{"2020-01-01 13:14:15", "-10m", "2020-01-01 13:04:15"},
		{"2020-01-01 13:14:15", "10.5m", "2020-01-01 13:24:45"},
		{"2020-01-01 13:14:15", "-10.5m", "2020-01-01 13:03:45"},

		{"2020-01-01 13:14:15", "10s", "2020-01-01 13:14:25"},
		{"2020-01-01 13:14:15", "-10s", "2020-01-01 13:14:05"},
		{"2020-01-01 13:14:15", "10.5s", "2020-01-01 13:14:25"},
		{"2020-01-01 13:14:15", "-10.5s", "2020-01-01 13:14:04"},

		{"2020-01-01 13:14:15", "-10x", ""},
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("catch an exception in Duration()：%s\n", r)
		}
	}()

	for _, v := range Tests {
		output := Parse(v.input).Duration(v.duration).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).Duration(v.duration).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddYears(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		years  int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2023-01-01"},
		{"2020-01-31", 3, "2023-01-31"},
		{"2020-02-01", 3, "2023-02-01"},
		{"2020-02-28", 3, "2023-02-28"},
		{"2020-02-29", 3, "2023-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).AddYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_NextYears(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		years  int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2023-01-01"},
		{"2020-01-31", 3, "2023-01-31"},
		{"2020-02-01", 3, "2023-02-01"},
		{"2020-02-28", 3, "2023-02-28"},
		{"2020-02-29", 3, "2023-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).NextYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).NextYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubYears(t *testing.T) {
	type Test struct {
		input  string // 输入值
		years  int    // 输入参数
		output string // 期望输出值
	}

	Tests := []Test{
		{"2020-01-01", 3, "2017-01-01"},
		{"2020-01-31", 3, "2017-01-31"},
		{"2020-02-01", 3, "2017-02-01"},
		{"2020-02-28", 3, "2017-02-28"},
		{"2020-02-29", 3, "2017-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).SubYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_PreYears(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		years  int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2017-01-01"},
		{"2020-01-31", 3, "2017-01-31"},
		{"2020-02-01", 3, "2017-02-01"},
		{"2020-02-28", 3, "2017-02-28"},
		{"2020-02-29", 3, "2017-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).PreYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).PreYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2021-01-01"},
		{"2020-01-31", "2021-01-31"},
		{"2020-02-01", "2021-02-01"},
		{"2020-02-28", "2021-02-28"},
		{"2020-02-29", "2021-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).AddYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_NextYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2021-01-01"},
		{"2020-01-31", "2021-01-31"},
		{"2020-02-01", "2021-02-01"},
		{"2020-02-28", "2021-02-28"},
		{"2020-02-29", "2021-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).NextYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).NextYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-01-01"},
		{"2020-01-31", "2019-01-31"},
		{"2020-02-01", "2019-02-01"},
		{"2020-02-28", "2019-02-28"},
		{"2020-02-29", "2019-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).SubYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_PreYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-01-01"},
		{"2020-01-31", "2019-01-31"},
		{"2020-02-01", "2019-02-01"},
		{"2020-02-28", "2019-02-28"},
		{"2020-02-29", "2019-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).PreYear().ToDateString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).PreYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMonths(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2020-04-01"},
		{"2020-01-31", 3, "2020-05-01"},
		{"2020-02-01", 3, "2020-05-01"},
		{"2020-02-28", 3, "2020-05-28"},
		{"2020-02-29", 3, "2020-05-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).AddMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_NextMonths(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2020-04-01"},
		{"2020-01-31", 3, "2020-04-30"},
		{"2020-02-01", 3, "2020-05-01"},
		{"2020-02-28", 3, "2020-05-28"},
		{"2020-02-29", 3, "2020-05-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).NextMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).NextMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMonths(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2019-10-01"},
		{"2020-01-31", 3, "2019-10-31"},
		{"2020-02-01", 3, "2019-11-01"},
		{"2020-02-28", 3, "2019-11-28"},
		{"2020-02-29", 3, "2019-11-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).SubMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_PreMonths(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2019-10-01"},
		{"2020-01-31", 3, "2019-10-31"},
		{"2020-02-01", 3, "2019-11-01"},
		{"2020-02-28", 3, "2019-11-28"},
		{"2020-02-29", 3, "2019-11-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).PreMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).PreMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2020-02-01"},
		{"2020-01-31", "2020-03-02"},
		{"2020-02-01", "2020-03-01"},
		{"2020-02-28", "2020-03-28"},
		{"2020-02-29", "2020-03-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).AddMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_NextMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2020-02-01"},
		{"2020-01-31", "2020-02-29"},
		{"2020-02-01", "2020-03-01"},
		{"2020-02-28", "2020-03-28"},
		{"2020-02-29", "2020-03-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).NextMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).NextMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-12-01"},
		{"2020-01-31", "2019-12-31"},
		{"2020-02-01", "2020-01-01"},
		{"2020-02-28", "2020-01-28"},
		{"2020-02-29", "2020-01-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).SubMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_PreMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-12-01"},
		{"2020-01-31", "2019-12-31"},
		{"2020-02-01", "2020-01-01"},
		{"2020-02-28", "2020-01-28"},
		{"2020-02-29", "2020-01-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).PreMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).PreMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddDays(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		days   int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2020-01-04"},
		{"2020-01-31", 3, "2020-02-03"},
		{"2020-02-01", 3, "2020-02-04"},
		{"2020-02-28", 3, "2020-03-02"},
		{"2020-02-29", 3, "2020-03-03"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddDays(v.days).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).AddDays(v.days).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubDays(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		days   int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2019-12-29"},
		{"2020-01-31", 3, "2020-01-28"},
		{"2020-02-01", 3, "2020-01-29"},
		{"2020-02-28", 3, "2020-02-25"},
		{"2020-02-29", 3, "2020-02-26"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubDays(v.days).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).SubDays(v.days).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2020-01-02"},
		{"2020-01-31", "2020-02-01"},
		{"2020-02-01", "2020-02-02"},
		{"2020-02-28", "2020-02-29"},
		{"2020-02-29", "2020-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddDay().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).AddDay().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-12-31"},
		{"2020-01-31", "2020-01-30"},
		{"2020-02-01", "2020-01-31"},
		{"2020-02-28", "2020-02-27"},
		{"2020-02-29", "2020-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubDay().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).SubDay().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddHours(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		hours  int
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 16:14:15"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 16:14:15"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 16:14:15"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 16:14:15"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 16:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddHours(v.hours).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).AddHours(v.hours).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubHours(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		hours  int
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 10:14:15"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 10:14:15"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 10:14:15"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 10:14:15"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 10:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubHours(v.hours).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).SubHours(v.hours).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 14:14:15"},
		{"2020-01-31 13:14:15", "2020-01-31 14:14:15"},
		{"2020-02-01 13:14:15", "2020-02-01 14:14:15"},
		{"2020-02-28 13:14:15", "2020-02-28 14:14:15"},
		{"2020-02-29 13:14:15", "2020-02-29 14:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddHour().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).AddHour().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 12:14:15"},
		{"2020-01-31 13:14:15", "2020-01-31 12:14:15"},
		{"2020-02-01 13:14:15", "2020-02-01 12:14:15"},
		{"2020-02-28 13:14:15", "2020-02-28 12:14:15"},
		{"2020-02-29 13:14:15", "2020-02-29 12:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubHour().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).SubHour().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMinutes(t *testing.T) {
	Tests := []struct {
		input   string // 输入值
		minutes int
		output  string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 13:17:15"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 13:17:15"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 13:17:15"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 13:17:15"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 13:17:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMinutes(v.minutes).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).AddMinutes(v.minutes).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMinutes(t *testing.T) {
	Tests := []struct {
		input   string // 输入值
		minutes int
		output  string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 13:11:15"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 13:11:15"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 13:11:15"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 13:11:15"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 13:11:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMinutes(v.minutes).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).SubMinutes(v.minutes).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 13:15:15"},
		{"2020-01-31 13:14:15", "2020-01-31 13:15:15"},
		{"2020-02-01 13:14:15", "2020-02-01 13:15:15"},
		{"2020-02-28 13:14:15", "2020-02-28 13:15:15"},
		{"2020-02-29 13:14:15", "2020-02-29 13:15:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMinute().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).AddMinute().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 13:13:15"},
		{"2020-01-31 13:14:15", "2020-01-31 13:13:15"},
		{"2020-02-01 13:14:15", "2020-02-01 13:13:15"},
		{"2020-02-28 13:14:15", "2020-02-28 13:13:15"},
		{"2020-02-29 13:14:15", "2020-02-29 13:13:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMinute().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).SubMinute().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddSeconds(t *testing.T) {
	Tests := []struct {
		input   string // 输入值
		seconds int
		output  string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 13:14:18"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 13:14:18"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 13:14:18"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 13:14:18"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 13:14:18"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddSeconds(v.seconds).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).AddSeconds(v.seconds).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubSeconds(t *testing.T) {
	Tests := []struct {
		input   string // 输入值
		seconds int
		output  string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 13:14:12"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 13:14:12"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 13:14:12"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 13:14:12"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 13:14:12"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubSeconds(v.seconds).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).SubSeconds(v.seconds).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddSecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 13:14:16"},
		{"2020-01-31 13:14:15", "2020-01-31 13:14:16"},
		{"2020-02-01 13:14:15", "2020-02-01 13:14:16"},
		{"2020-02-28 13:14:15", "2020-02-28 13:14:16"},
		{"2020-02-29 13:14:15", "2020-02-29 13:14:16"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddSecond().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).AddSecond().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubSecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 13:14:14"},
		{"2020-01-31 13:14:15", "2020-01-31 13:14:14"},
		{"2020-02-01 13:14:15", "2020-02-01 13:14:14"},
		{"2020-02-28 13:14:15", "2020-02-28 13:14:14"},
		{"2020-02-29 13:14:15", "2020-02-29 13:14:14"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubSecond().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := Timezone(PRC).Parse(v.input).SubSecond().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}
