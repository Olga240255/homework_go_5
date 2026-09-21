package common

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// TODO: ReadFileOrDefault должен прочитать файл. Если файла нет — вернуть defaultValue без ошибки. Другие ошибки вернуть наружу.
func ReadFileOrDefault(path string, defaultValue string) (string, error) {
	J, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return defaultValue, nil
		} else {
			return "", err
		}
	}
	defer J.Close()
	s, err := io.ReadAll(J)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

// TODO: WriteFileLines должен записать строки в файл, каждую с переводом строки в конце.
func WriteFileLines(path string, lines []string) error {
	F, err := os.Create(path)
	if err != nil {
		return err
	}
	defer F.Close()
	for _, val := range lines {
		_, err = F.Write([]byte(val))
		if err != nil {
			return err
		}
		_, err = F.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}
	if err := F.Sync(); err != nil {
		return err
	}

	return nil
}

// TODO: CountFileLines должен открыть файл, посчитать строки через Scanner и вернуть ошибку открытия или сканирования.
func CountFileLines(path string) (int, error) {
	F, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer F.Close()
	i := 0
	s := bufio.NewScanner(F)
	for s.Scan() {
		i++
	}
	if err = s.Err(); err != nil {
		return 0, err
	}
	return i, nil
}

func Example() string {
	var out strings.Builder
	out.WriteString("common file tasks")
	return out.String()
}
