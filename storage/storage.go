package storage

import (
	"fmt"
	"github.com/koyeo/snippet/logger"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Abs(path ...string) string {
	return filepath.Join(strings.TrimPrefix(filepath.Join(path[:]...), Root()))
}

func Root() string {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return root
}

func Join(paths ...string) string {

	for i, v := range paths {
		paths[i] = v
	}

	return Abs(paths...)
}

func ReadFile(path string) (content string, err error) {

	if !PathExist(path) {
		return
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	content = string(data)

	return
}

func ReadFiles(debug bool, path string, ignore []string, prefix []string, suffix []string) (files []string, err error) {

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	sep := string(os.PathSeparator)

	for _, fi := range dir {
		path := filepath.Join(path, sep, fi.Name())
		items := strings.Split(path, sep)
		subPath := strings.Join(items[1:], sep)
		var ok bool
		for _, v := range ignore {
			ok, err = filepath.Match(v, subPath)
			if err != nil {
				logger.Fatal("Match ignore path error: ", err)
			}
			if ok {
				break
			}
		}
		if ok {
			if debug {
				logger.IgnoreReadPath("Read files ignore path:", subPath)
			}
			continue
		}
		if fi.IsDir() {
			var items []string
			items, err = ReadFiles(debug, path, ignore, prefix, suffix, )
			if err != nil {
				return
			}
			files = append(files, items...)
		} else {
			if HasPrefix(fi.Name(), prefix...) || HasSuffix(fi.Name(), suffix...) {
				files = append(files, path)
			}
		}
	}

	return
}

func ReadDirs(debug bool, path string, ignore []string, prefix []string, suffix []string) (dirs []string, err error) {

	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	sep := string(os.PathSeparator)

	for _, fi := range infos {
		path := filepath.Join(path, sep, fi.Name())
		items := strings.Split(path, sep)
		subPath := strings.Join(items[1:], sep)
		var ok bool
		for _, v := range ignore {
			ok, err = filepath.Match(v, subPath)
			if err != nil {
				logger.Fatal("Match ignore path error: ", err)
			}
			if ok {
				break
			}
		}
		if ok {
			if debug {
				logger.IgnoreReadPath("Read dirs ignore path:", subPath)
			}
			continue
		}

		if fi.IsDir() {
			if HasPrefix(fi.Name(), prefix...) || HasSuffix(fi.Name(), suffix...) {
				dirs = append(dirs, path)
			}
			var items []string
			items, err = ReadDirs(debug, path, ignore, prefix, suffix)

			dirs = append(dirs, items...)
		}
	}

	return
}

func HasPrefix(name string, prefix ...string) bool {

	for _, v := range prefix {
		v = strings.TrimSpace(v)
		if v != "" && strings.HasPrefix(name, v) {
			return true
		}
	}

	return false
}

func HasSuffix(name string, suffix ...string) bool {
	for _, v := range suffix {
		v = strings.TrimSpace(v)
		if v != "" && strings.HasSuffix(name, v) {
			return true
		}
	}

	return false
}

func MakeDir(path string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		cmd := exec.Command("bash", "-c", fmt.Sprintf("mkdir -p %s", path))
		_ = cmd.Run()
		logger.MakeDirSuccess(path)
	}

}

func PathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else if err != nil {
		logger.Fatal("Check file exists:", err)
	}
	return true
}

func Remove(path string) (err error) {
	err = os.RemoveAll(path)
	return
}

func WriteFile(path string, content []byte) (err error) {
	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return
	}
	return
}

func MakePath(path, makePrefix, name, makeSuffix, suffix string) string {
	customPath := filepath.Join(path, name, suffix)
	if PathExist(customPath) {
		return customPath
	}
	return filepath.Join(path, makePrefix, name, makeSuffix, suffix)
}
