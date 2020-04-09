package storage

import (
	"fmt"
	"github.com/koyeo/snippet/logger"
	"io/ioutil"
	"os"
	"os/exec"
	goPath "path"
	"path/filepath"
	"regexp"
	"strings"
)

func Join(paths ...string) string {

	for i, v := range paths {
		paths[i] = v
	}

	return Abs(paths...)
}

func ReadFile(path string) (content string, err error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	content = string(data)

	return
}

func Remove(path string) (err error) {
	err = os.RemoveAll(path)
	return
}

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else if err != nil {
		logger.Fatal("Check file exists:", err)
	}
	return true
}

func Files(path string, suffix ...string) (files []string, err error) {

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	sep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() {
			var dirFiles []string
			dirFiles, err = Files(filepath.Join(path, sep, fi.Name()), suffix...)
			if err != nil {
				return
			}
			files = append(files, dirFiles...)
		} else {

			if hasSuffix(suffix, fi.Name()) {
				files = append(files, filepath.Join(path, sep, fi.Name()))
			}
		}
	}

	return
}

func hasSuffix(suffix []string, fileName string) bool {
	for _, v := range suffix {
		if strings.HasSuffix(fileName, v) {
			return true
		}
	}

	return false
}

func Root() string {

	root, err := os.Getwd()
	if err != nil {
		logger.Fatal("Get root path error: ", err)
	}

	return root
}

func MakeDir(path string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		// os.Mkdir(path, os.ModePerm)
		cmd := exec.Command("bash", "-c", fmt.Sprintf("mkdir -p %s", path))
		_ = cmd.Run()
		logger.MakeDirSuccess(path)
	}

}

func MakeDirQuiet(path string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		cmd := exec.Command("bash", "-c", fmt.Sprintf("mkdir -p %s", path))
		_ = cmd.Run()
	}

}

func Abs(path ...string) string {
	return filepath.Join(strings.TrimPrefix(filepath.Join(path[:]...), Root()))
}

func WriteFile(path string, content []byte) (err error) {
	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return
	}
	return
}

func RelativePath(path string) string {
	wd, _ := os.Getwd()
	path = strings.TrimPrefix(path, wd+"/")
	return path
}

func PackagePath(path string) string {
	wd, _ := os.Getwd()
	items := strings.Split(wd, "/")

	if len(items) == 0 {
		return path
	}

	return filepath.Join(items[len(items)-1], path)
}

func ParseFilePath(filePath string) (path, file, ext string) {
	file = goPath.Base(filePath)
	path = strings.TrimSuffix(filePath, file)
	path = strings.TrimRight(path, "/")
	ext = goPath.Ext(file)
	file = strings.TrimSuffix(file, ext)
	ext = strings.TrimLeft(ext, ".")
	return
}

func ParseYamlMapLines(content string, fields ...string) (lines []string) {
	items := strings.Split(content, "\n")
	if len(fields) == 0 {
		lines = items
		return
	}

	depth := 0
	expected := len(fields)
	spacesRegex := regexp.MustCompile(`^(\s+).+\s*$`)

	index := 0
	field := fmt.Sprintf("%s:", fields[index])
	lineSpacesCount := 0

	for l, v := range items {

		if depth != expected {
			if strings.TrimSpace(v) == field {
				if index < expected-1 {
					index++
					field = fmt.Sprintf("%s:", fields[index])
				}

				depth++
				if depth == expected {
					continue
				}
			}

		} else {

			strs := spacesRegex.FindStringSubmatch(v)

			if len(strs) != 2 {
				break
			}

			if lineSpacesCount == 0 {
				lineSpacesCount = len(strs[1])
			}

			if len(strs[1]) != lineSpacesCount {
				break
			}
			lines = append(lines, strings.TrimSpace(v))
		}

		if depth != 0 {
			if l+1 < len(items)-1 {
				strs := spacesRegex.FindStringSubmatch(items[l+1])
				if len(strs) == 0 {
					break
				}
			}
		}
	}

	return
}
