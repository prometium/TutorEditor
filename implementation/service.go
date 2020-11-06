package implementation

import (
	"archive/zip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"editorsvc"
	"editorsvc/utils"
)

type service struct {
	repository editorsvc.Repository
}

// NewService makes a new Service
func NewService(rep editorsvc.Repository) editorsvc.Service {
	return &service{
		repository: rep,
	}
}

func (s *service) Setup(ctx context.Context) error {
	return s.repository.Setup(ctx)
}

func (s *service) AddRawScript(ctx context.Context, name string, fileReader io.Reader) (string, error) {
	exists, err := s.repository.ScriptExists(ctx, name)
	if err != nil {
		return "", err
	} else if exists {
		name = fmt.Sprintf("%s [%s]", name, time.Now().Format("2006-01-02 15:04:05"))
	}

	zipReader, err := utils.CreateZipReader(fileReader)
	if err != nil {
		return "", err
	}

	imagesDir := fmt.Sprintf("static/images/")
	os.MkdirAll(imagesDir, os.ModePerm)

	var linksMap map[string]string = make(map[string]string)
	var scriptFile *zip.File = nil
	for _, file := range zipReader.File {
		if filepath.Ext(strings.TrimSpace(file.Name)) == ".png" {
			hash, err := utils.HashZipFileMD5(file)
			if err != nil {
				return "", err
			}
			linksMap[file.Name] = hash + ".png"
			// err = utils.CopyZipFile(file, filepath.Join(imagesDir, linksMap[file.Name]))
			// if err != nil {
			// 	return "", err
			// }
		} else if file.Name == "Script.json" {
			scriptFile = file
		}
	}

	frames, err := getFramesFromScriptFile(scriptFile)
	if err != nil {
		return "", err
	}

	configuredFrames := configurePictureLinks(frames, "images/", linksMap)
	id, err := s.repository.AddScript(ctx, name, configuredFrames)
	if err != nil {
		return "", err
	}
	return id, nil
}

func getFramesFromScriptFile(scriptFile *zip.File) ([]editorsvc.Frame, error) {
	scriptJSON, err := utils.ReadAllFromZip(scriptFile)
	if err != nil {
		return nil, err
	}

	var rs rawScript
	if err := json.Unmarshal([]byte(scriptJSON), &rs); err != nil {
		return nil, err
	}

	frames := make([]editorsvc.Frame, len(rs.Frames))
	for i, frame := range rs.Frames {
		frames[i] = editorsvc.Frame{
			UID:         strconv.Itoa(frame.FrameNumber),
			PictureLink: frame.PictureLink,
			Task: editorsvc.Task{
				Text: frame.Task,
			},
			Hint: editorsvc.Hint{
				Text: frame.Hint,
			},
		}
	}

	for i, frame := range rs.Frames[1:] {
		action := &frame.ActionSwitch
		frames[i].Actions = []editorsvc.Action{
			editorsvc.Action{
				NextFrame: editorsvc.NextFrame{
					UID: strconv.Itoa(frame.FrameNumber),
				},
				ActionType:   action.ActionType,
				XLeft:        action.XLeft,
				XRight:       action.XRight,
				YLeft:        action.YLeft,
				YRight:       action.YRight,
				StartXLeft:   action.StartXLeft,
				StartYLeft:   action.StartYLeft,
				StartXRight:  action.StartXRight,
				StartYRight:  action.StartYRight,
				FinishXLeft:  action.FinishXLeft,
				FinishYLeft:  action.FinishYLeft,
				FinishXRight: action.FinishXRight,
				FinishYRight: action.FinishYRight,
				TicksCount:   action.TicksCount,
				Key:          action.Key,
				ModKey:       action.ModKey,
			},
		}
	}

	return frames, nil
}

func configurePictureLinks(frames []editorsvc.Frame, path string, linksMap map[string]string) []editorsvc.Frame {
	for i := range frames {
		frames[i].PictureLink = filepath.Join(path, linksMap[frames[i].PictureLink])
	}
	return frames
}

func (s *service) GetScriptsList(ctx context.Context) ([]editorsvc.Script, error) {
	return s.repository.GetScriptsList(ctx)
}
