package snippet

import (
	"fmt"
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

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content = string(data)

	return
}

func ReadFiles(path string, prefix []string, suffix []string) (files []string, err error) {

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	sep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() {
			var items []string
			items, err = ReadFiles(filepath.Join(path, sep, fi.Name()), prefix, suffix, )
			if err != nil {
				return
			}
			files = append(files, items...)
		} else {
			if HasPrefix(fi.Name(), prefix...) || HasSuffix(fi.Name(), suffix...) {
				files = append(files, filepath.Join(path, sep, fi.Name()))
			}
		}
	}

	return
}

func ReadDirs(root string, prefix []string, suffix []string) (dirs []string, err error) {

	infos, err := ioutil.ReadDir(root)
	if err != nil {
		return
	}

	sep := string(os.PathSeparator)

	for _, fi := range infos {
		if fi.IsDir() {
			if HasPrefix(fi.Name(), prefix...) || HasSuffix(fi.Name(), suffix...) {
				dirs = append(dirs, filepath.Join(root, sep, fi.Name()))
			}
			var items []string
			items, err = ReadDirs(filepath.Join(root, sep, fi.Name()), prefix, suffix, )

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
		MakeDirSuccess(path)
	}

}

func PathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else if err != nil {
		Fatal("Check file exists:", err)
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
