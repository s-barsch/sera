package video

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"g.rg-s.com/sacer/go/entry"
	"g.rg-s.com/sacer/go/entry/file"
	"g.rg-s.com/sacer/go/entry/info"
	"g.rg-s.com/sacer/go/entry/tools"
	"g.rg-s.com/sacer/go/entry/tools/script"
	"g.rg-s.com/sacer/go/entry/tools/transcript"
	"g.rg-s.com/sacer/go/requests/split"

	"github.com/alfg/mp4"
)

type Video struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	Sources []*Source

	Captions   []string
	Transcript *script.Script

	Duration float64
}

type Source struct {
	Path       string
	Size       int64
	Resolution string
}

func (s *Source) Mbyte() int64 {
	return s.Size / 1024 / 1024
}

var resWidth = map[string]string{
	"1080": "1920",
	"720":  "1280",
	"480":  "854",
}

func (s *Source) ResolutionWidth() string {
	return resWidth[s.Resolution]
}

func NewVideo(path string, parent entry.Entry) (*Video, error) {
	fnErr := &tools.Err{
		Path: path,
		Func: "NewVideo",
	}

	file, err := file.NewFile(path)
	if err != nil {
		return nil, err
	}

	sources, err := getSources(path)
	if err != nil {
		return nil, err
	}

	duration, err := Mp4Duration(path)
	if err != nil {
		return nil, err
	}

	inf := info.Info{}

	if info.HasFileInfo(path) {
		i, err := info.ReadFileInfo(path)
		if err != nil {
			fnErr.Err = err
			return nil, fnErr
		}
		inf = i
	}

	date, err := tools.ParseTimestamp(inf["date"])
	if err != nil {
		date, err = tools.ParseDatePath(path)
		if err != nil {
			fnErr.Err = err
			return nil, fnErr
		}
	}

	captions := getCaptions(path)

	script, err := transcript.GetTranscripts(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	return &Video{
		parent:     parent,
		file:       file,
		date:       date,
		info:       inf,
		Captions:   captions,
		Transcript: script,
		Sources:    sources,
		Duration:   duration,
	}, nil
}

func getCaptions(path string) []string {
	dir := filepath.Dir(path)
	name := stripResolution(tools.StripExt(filepath.Base(path)))
	langs := []string{}
	for _, lang := range []string{"de", "en"} {
		_, err := os.Stat(filepath.Join(dir, "vtt", fmt.Sprintf("%v.%v.vtt", name, lang)))
		if err == nil {
			langs = append(langs, lang)
		}
	}
	return langs
}

func getSources(path string) ([]*Source, error) {

	top, _ := getSource(path)

	sources := []*Source{top}

	ress := filepath.Dir(path) + "/sizes"
	_, err := os.Stat(ress)
	if err != nil {
		return sources, nil
	}

	l, err := os.ReadDir(ress)
	if err != nil {
		return nil, err
	}

	sort.Sort(Desc(l))

	dir := filepath.Dir(path)

	for _, fi := range l {
		if fi.IsDir() || tools.IsNameSys(fi.Name()) {
			continue
		}

		s, err := getSource(filepath.Join(dir, "sizes", fi.Name()))
		if err != nil {
			return nil, err
		}

		if s.Resolution == "1080" && sources[0].Resolution == "1080" {
			sources[0] = s
			continue
		}

		sources = append(sources, s)
	}
	return sources, nil
}

func parent(path string) string {
	return filepath.Base(filepath.Dir(path))
}

func getSource(path string) (*Source, error) {
	isTop := false
	name := ""

	if parent(path) == "sizes" {
		name = "sizes/" + filepath.Base(path)
	} else {
		name = filepath.Base(path)
		isTop = true
	}

	file, err := split.SplitFile(path)
	if err != nil {
		return nil, fmt.Errorf("getSource: cannot find source since path is malformed. %w", err)
	}

	res, err := strconv.Atoi(file.Option)
	if err != nil {
		if isTop {
			res = 1080
		} else {
			return nil, fmt.Errorf("getSources: Could not find resolution of %v", path)
		}
	}
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &Source{
		Path:       name,
		Size:       fi.Size(),
		Resolution: strconv.Itoa(res),
	}, nil
}

func stripResolution(name string) string {
	i := strings.LastIndex(name, "-")
	if i > 0 {
		return name[:i]
	}
	return name
}

type Desc []os.DirEntry

func (a Desc) Len() int      { return len(a) }
func (a Desc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a Desc) Less(i, j int) bool {
	return a[i].Name() > a[j].Name()
}

/*
func getDuration(path string) (uint32, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	m, err := mp4.Decode(io.Reader(f))
	if err != nil {
		println(path)
		println("here")
		return 0, err
	}
	return m.Moov.Mvhd.Duration, nil
}
*/

func Mp4Duration(path string) (float64, error) {
	f, err := mp4.Open(path)
	if err != nil {
		return 0, err
	}
	return float64(f.Moov.Mvhd.Duration / 1000), nil
}
