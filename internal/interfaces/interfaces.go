package interfaces

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	ID   int
	Name string
}

// TODO: String должен вернуть пользователя в формате "user#<ID>:<Name>".
func (u User) String() string {
	return "user#" + strconv.Itoa(u.ID) + ":" + u.Name
}

// TODO: FormatOne должен вызвать String у переданного значения. Если значение nil, верни пустую строку.
func FormatOne(value fmt.Stringer) string {
	if value == nil {
		return ""
	}
	return value.String()
}

// TODO: FormatMany должен вернуть строки для всех элементов в исходном порядке. nil элементы пропускаются.
func FormatMany(values []fmt.Stringer) []string {
	res := make([]string, 0)
	if values != nil {
		for _, val := range values {
			if val != nil {
				res = append(res, val.String())
			}
		}
	}
	return res
}

type EmailNotifier struct {
	Address string
}

// TODO: Notify должен вернуть "email:<Address>:<message>".
func (n EmailNotifier) Notify(message string) string {
	return "email:" + n.Address + ":" + message
}

type PushNotifier struct {
	DeviceID string
}

// TODO: Notify должен вернуть "push:<DeviceID>:<message>".
func (n PushNotifier) Notify(message string) string {
	return "push:" + n.DeviceID + ":" + message
}

type Notifier interface {
	Notify(message string) string
}

// TODO: SendNotification должен отправить одно сообщение через Notifier.
// Если Notifier nil, верни пустую строку.
func SendNotification(n Notifier, message string) string {
	if n == nil {
		return ""
	}
	return n.Notify(message)
}

// TODO: SendBatch должен отправить все сообщения и вернуть результаты в
// том же порядке. nil Notifier даёт пустой результат.
func SendBatch(n Notifier, messages []string) []string {
	if n == nil {
		return []string{}
	}
	a := make([]string, 0)
	for _, val := range messages {
		a = append(a, n.Notify(val))
	}
	return a
}

type StaticLoader struct {
	Value string
	Err   error
}

// TODO: Load должен вернуть поля Value и Err без дополнительной обработки.
func (l StaticLoader) Load() (string, error) {
	return l.Value, l.Err
}

type Loader interface {
	Load() (string, error)
}

// TODO: LoadUpper должен загрузить строку, привести её к верхнему регистру
//
//	и вернуть ошибку без потери, если загрузка не удалась.
func LoadUpper(l Loader) (string, error) {
	if l == nil {
		return "", errors.New("false")
	}
	s, err := l.Load()
	if err != nil {
		return "", err
	}
	return strings.ToUpper(s), err
}

// TODO: Describe должен вернуть описание динамического типа: nil, string:<value>,
// user:<String>, notifier:<Notify("ping")>, loader:<loaded>, error:<text>
// или unknown:<T>.
func Describe(value any) string {
	if value == nil {
		return "nil"
	}
	switch v := value.(type) {
	case User:
		{
			return fmt.Sprintf("user:%s", v.String())
		}
	case string:
		{
			return fmt.Sprintf("string:%s", v)
		}
	case Notifier:
		{
			return fmt.Sprintf("notifier:%s", v.Notify("ping"))
		}
	case Loader:
		{
			s, e := v.Load()
			if e == nil {
				return fmt.Sprintf("loader:%s", s)
			} else {
				return fmt.Sprintf("loader-error:%s", e.Error())
			}
		}
	case error:
		{
			return fmt.Sprintf("error:%s", v.Error())
		}
	default:
		{
			return fmt.Sprintf("unknown:%T", v)
		}

	}
}

// TODO: OnlyErrors должен вернуть тексты только ненулевых ошибок в исходном порядке.
func OnlyErrors(values []error) []string {
	res := make([]string, 0)
	if values == nil {
		return res
	}
	for _, val := range values {
		if val != nil {
			res = append(res, val.Error())
		}
	}
	return res
}

// TODO: JoinStringers должен отформатировать ненулевые Stringer и объединить их через sep.
func JoinStringers(values []fmt.Stringer, sep string) string {
	if values == nil {
		return ""
	}
	var s string
	t := false
	for _, val := range values {
		if val != nil {
			if t == false {
				s = val.String()
				t = true
			} else {
				s = s + sep + val.String()
			}

		}
	}
	return s
}

// TODO: FirstNonEmptyStringer должен вернуть первую непустую строку среди Stringer. nil и пустые строки пропускаются.
func FirstNonEmptyStringer(values []fmt.Stringer) string {
	if values == nil {
		return ""
	}
	var s string
	for _, val := range values {
		if val != nil {
			if val.String() != "" {
				s = val.String()
				break
			}
		}
	}
	return s
}

// TODO: CountStringers должен посчитать, сколько элементов реально реализуют fmt.Stringer.
func CountStringers(values []any) int {
	if values == nil {
		return 0
	}
	var count int
	for _, val := range values {
		_, ok := val.(fmt.Stringer)
		if ok == true {
			count++
		}
	}
	return count
}

// TODO: NormalizeAndFormat должен обрезать пробелы вокруг String() и схлопнуть пустой результат в "empty".
func NormalizeAndFormat(value fmt.Stringer) string {
	if value == nil {
		return "empty"

	}
	s := strings.TrimSpace(value.String())
	if s == "" {
		return "empty"
	}
	return s
}

func Example() string {
	var out strings.Builder
	fmt.Fprintf(&out, "user: %s\n", FormatOne(User{ID: 7, Name: "Maria"}))
	fmt.Fprintf(&out, "email: %s\n", SendNotification(EmailNotifier{Address: "team@example.com"}, "hello"))
	fmt.Fprintf(&out, "push: %s\n", SendNotification(PushNotifier{DeviceID: "ios-1"}, "hello"))
	fmt.Fprintf(&out, "describe: %s", Describe(User{ID: 8, Name: "Alex"}))
	return out.String()
}
