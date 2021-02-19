package internal

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	// 从字符串 "#INCLUDE [dir1/dir2/file.txt] @SECTION_1" 中提取出 "dir1/dir2/file.txt" 和 "@SECTION_1"
	RegexInclude = regexp.MustCompile(`#INCLUDE\s*\[([^\n]+)\]\s*(@[^\n]+)`)

	// 从字符串 "#INSERT [dir1/dir2/file.txt]" 中提取出 "dir1/dir2/file.txt"
	RegexInsert = regexp.MustCompile(`#INSERT\s*\[([^\n]+)\]\s*$`)

	// 从字符串 "[@MAIN]" 中提取出 "@MAIN"
	RegexFunctionName = regexp.MustCompile(`^\[([^\n]+)\]\s*$`)

	// 从字符串 "(PARAM)" 中提取出 "PARAM"
	RegexFunctionParam = regexp.MustCompile(`\(([\S]+)\)`)

	// 从字符串 "<$PARAM>" 中提取出 "$PARAM"
	RegexFormatParam = regexp.MustCompile(`\<(\$\w+)\>`)
)

func LoadFile(file string) (*Script, error) {
	var r, err = os.Open(file)
	if r != nil {
		defer r.Close()
	}
	if err != nil {
		return nil, err
	}
	return Load(r)
}

func Load(r io.Reader) (*Script, error) {
	var lines, err = ExpandScript(Read(r))
	if err != nil {
		return nil, err
	}

	var s *Section
	var script = NewScript()

	for _, line := range lines {
		if line[0] == '[' {
			var match = RegexFunctionName.FindStringSubmatch(line)
			if len(match) > 0 {
				if s != nil {
					script.Add(s)
				}

				s = NewSection(match[1])
				continue
			}
		}

		if s != nil {
			s.Add(line)
		}
	}

	if s != nil {
		script.Add(s)
	}

	return script, nil
}

// ExpandScript 处理 #INSERT 语句和 #INCLUDE 语句
func ExpandScript(lines []string) ([]string, error) {
	var nLines []string
	for _, line := range lines {
		if SkipLine(line) {
			continue
		}

		if line[0] == KeyPrefix {
			if strings.HasPrefix(line, KeyInsert) {
				var match = RegexInsert.FindStringSubmatch(line)
				var insertLines, err = ReadFile(match[1])
				if err != nil {
					return nil, err
				}
				insertLines, err = ExpandScript(insertLines)
				if err != nil {
					return nil, err
				}
				nLines = append(nLines, insertLines...)
				continue
			} else if strings.HasPrefix(line, KeyInclude) {
				var match = RegexInclude.FindStringSubmatch(line)
				var insertLines, err = Include(match[1], match[2])
				if err != nil {
					return nil, err
				}
				nLines = append(nLines, insertLines...)
				continue
			}
		}
		nLines = append(nLines, line)
	}
	return nLines, nil
}

func ReadFile(file string) ([]string, error) {
	var r, err = os.Open(file)
	if r != nil {
		defer r.Close()
	}
	if err != nil {
		return nil, err
	}
	return Read(r), nil
}

func Read(r io.Reader) []string {
	var lines []string

	var scanner = bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, string(RemoveBOM(scanner.Bytes())))
	}
	return lines
}

// Include 读取指定文件中的指定片断
func Include(file, name string) ([]string, error) {
	var lines, err = ReadFile(file)
	if err != nil {
		return nil, err
	}

	name = "[" + name + "]"

	var stat = 0

	var nLines []string
	for _, line := range lines {
		if SkipLine(line) {
			continue
		}
		switch stat {
		case 0:
			if line[0] == '[' && strings.HasPrefix(line, name) {
				stat = 1
			}
		case 1:
			if line[0] == '[' {
				return nLines, nil
			}

			nLines = append(nLines, line)
		}
	}
	return nLines, nil
	//return nil, errors.New("syntax error:" + file)
}
func SkipLine(line string) bool {
	if line == "" {
		return true
	}
	if line[0] == KeyComment {
		return true
	}
	return false
}
