package ioflow

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

// TODO: ReadAllText должен прочитать все данные из io.Reader и вернуть строку.
func ReadAllText(r io.Reader) (string, error) {
	p, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(p), nil
}

// TODO: CountBytes должен посчитать количество байт, доступных через Reader, и вернуть ошибку чтения без потери.
func CountBytes(r io.Reader) (int, error) {
	n, err := io.Copy(io.Discard, r)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

// TODO: WriteString должен записать text в Writer и вернуть количество записанных байт и ошибку.
func WriteString(w io.Writer, text string) (int, error) {
	b := []byte(text)
	n, err := w.Write(b)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// TODO: WriteLines должен записать каждую строку с переводом строки. При ошибке остановиться и вернуть её.
func WriteLines(w io.Writer, lines []string) error {
	for _, val := range lines {
		b := append([]byte(val), byte('\n'))
		_, err := w.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: CopyAll должен скопировать все данные из src в dst и вернуть число байт.
func CopyAll(dst io.Writer, src io.Reader) (int64, error) {
	n, err := io.Copy(dst, src)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// TODO: CopyUpper должен прочитать src, привести текст к верхнему регистру и записать в dst.
func CopyUpper(dst io.Writer, src io.Reader) (int64, error) {
	b, err := io.ReadAll(src)
	if err != nil {
		return 0, err
	}
	s := []byte(strings.ToUpper(string(b)))
	n, err := dst.Write(s)
	if err != nil {
		return 0, err
	}
	return int64(n), nil
}

// TODO: ScanLines должен вернуть все строки из Reader без символов перевода строки.
func ScanLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	result := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

// TODO: CountNonEmptyLines должен посчитать строки, которые после TrimSpace не пустые.
func CountNonEmptyLines(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	res := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			res++
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return res, nil
}

// TODO: FirstLine должен вернуть первую строку. Если строк нет, вернуть пустую строку без ошибки.
func FirstLine(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	line := scanner.Text()
	return line, nil
}

// TODO: ReadCSVLike должен прочитать строки, разделить каждую по запятой и обрезать пробелы у ячеек.
func ReadCSVLike(r io.Reader) ([][]string, error) {
	scanner := bufio.NewScanner(r)
	res := make([][]string, 0)
	for scanner.Scan() {
		sub := strings.Split(scanner.Text(), ",")
		for i := range sub {
			sub[i] = strings.TrimSpace(sub[i])
		}
		res = append(res, sub)
	}
	if err := scanner.Err(); err != nil {
		return res, err
	}
	return res, nil
}

// TODO: LimitRead должен прочитать не больше limit байт и вернуть строку. Отрицательный limit считать нулём.
func LimitRead(r io.Reader, limit int) (string, error) {
	if limit <= 0 {
		return "", nil
	}
	limitedReader := io.LimitReader(r, int64(limit))
	bytes, err := io.ReadAll(limitedReader)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// TODO: RepeatToWriter должен записать text count раз подряд. count <= 0 ничего не пишет.
func RepeatToWriter(w io.Writer, text string, count int) error {
	if count <= 0 {
		return nil
	}
	for i := 0; i < count; i++ {
		_, err := io.WriteString(w, text)
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: BufferReport должен вернуть строки, записанные через Writer в формате "<index>:<line>\n".
func BufferReport(lines []string) string {
	if lines == nil || len(lines) == 0 {
		return ""
	}
	var b strings.Builder
	for i, val := range lines {
		b.WriteString(strconv.Itoa(i))
		b.WriteString(":")
		b.WriteString(val)
		b.WriteString("\n")
	}
	return b.String()
}

// TODO: ReadAndTrim должен прочитать весь Reader и вернуть строку без пробелов по краям.
func ReadAndTrim(r io.Reader) (string, error) {
	s, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(s)), nil
}

// TODO: WriteKeyValues должен записать пары map в отсортированном по ключу порядке: "key=value\n".
func WriteKeyValues(w io.Writer, values map[string]string) error {
	if values == nil || len(values) == 0 {
		return nil
	}
	s := make([]string, 0)
	for i := range values {
		s = append(s, i)
	}
	slices.Sort(s)
	for _, k := range s {
		_, err := fmt.Fprintf(w, "%s=%s\n", k, values[k])
		if err != nil {
			return err
		}
	}
	return nil
}

func Example() string {
	var out strings.Builder
	text, _ := ReadAllText(strings.NewReader("hello"))
	out.WriteString("read: " + text + "\n")
	count, _ := CountNonEmptyLines(strings.NewReader("a\n\n b \n"))
	out.WriteString("non-empty: ")
	out.WriteString(strings.TrimSpace(BufferReport([]string{string(rune('0' + count))})))
	return out.String()
}
