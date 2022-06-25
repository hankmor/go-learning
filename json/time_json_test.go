package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 自定义类型
type Time time.Time

const (
	// 时间默认格式
	timeFormat = "2006-01-02 15:04:05"
)

// 实现json的 UnmarshalJSON 和 MarshalJSON 方法自定义json序列化/反序列化
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

type Person struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Birthday Time   `json:"birthday"`
}

// 格式化为时间戳
type Time1 time.Time

func (t *Time1) UnmarshalJSON(data []byte) (err error) {
	// now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	tm := bytesToInt(data)
	*t = Time1(time.UnixMilli(int64(tm)))
	return
}

func (t Time1) MarshalJSON() ([]byte, error) {
	tm := time.Time(t)
	r := tm.UnixMilli()
	return intToBytes(int(r)), nil
}

func (t Time1) String() string {
	return strconv.FormatInt(time.Time(t).UnixMilli(), 10)
}

func intToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func bytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

type Person1 struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Birthday Time1  `json:"birthday"`
}

func TestTimeJson(t *testing.T) {
	now := Time(time.Now())
	t.Log(now)
	src := `{"id":5,"name":"xiaoming","birthday":"2016-06-30 16:09:51"}`
	p := new(Person)
	err := json.Unmarshal([]byte(src), p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
	t.Log(time.Time(p.Birthday))
	js, _ := json.Marshal(p)
	t.Log(string(js))

	fmt.Println("=======")
	t.Log(time.Now().UnixMilli())
	now1 := Time1(time.Now())
	t.Log(now1)

	p1 := Person1{
		Id:       5,
		Name:     "huzhou",
		Birthday: now1,
	}
	s, _ := json.Marshal(p1)
	_ = json.Unmarshal(s, &p1)
	t.Log(p1)
}
