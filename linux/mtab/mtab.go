package mtab

import (
	"bufio"
	"os"
	"strings"
)

type MountInfo struct {
	Source  string
	Target  string
	Type    string
	Options []string
}

type Mtab struct {
	Mounts []*MountInfo
}

func Open(filepath string) (*Mtab, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	mtab := &Mtab{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(strings.Trim(scanner.Text(), " \n"), " ")
		if len(line) != 6 {
			continue
		}
		mtab.Mounts = append(mtab.Mounts, &MountInfo{
			Source:  line[0],
			Target:  line[1],
			Type:    line[2],
			Options: strings.Split(line[3], ","),
		})
	}

	return mtab, nil
}

func (m *Mtab) SearchByTarget(dir string) (*MountInfo, bool) {
	return m.SearchBy(func(mi *MountInfo) bool {
		return mi.Target == dir
	})
}

func (m *Mtab) SearchBySource(source string) (*MountInfo, bool) {
	return m.SearchBy(func(mi *MountInfo) bool {
		return mi.Source == source
	})
}

func (m *Mtab) SearchBy(callback func(*MountInfo) bool) (*MountInfo, bool) {
	for _, mountInfo := range m.Mounts {
		if callback(mountInfo) {
			return mountInfo, true
		}
	}
	return nil, false
}

func (m *MountInfo) HasOption(opt string) (string, bool) {
	for _, mountOpt := range m.Options {
		if opt == mountOpt {
			return opt, true
		}
		if strings.HasPrefix(mountOpt, opt+"=") {
			return mountOpt[len(opt)+1:], true
		}
	}
	return "", false
}
