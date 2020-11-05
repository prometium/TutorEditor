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

func (s *service) AddRawScript(ctx context.Context, name string, archiveReader io.Reader) (string, error) {
	exists, err := s.repository.ScriptExists(ctx, name)
	if err != nil {
		return "", err
	} else if exists {
		return "", editorsvc.ErrScriptAlreadyExists
	}

	zipReader, err := utils.CreateZipReader(archiveReader)
	if err != nil {
		return "", err
	}

	dir := fmt.Sprintf("static/images/%s/", name)
	os.MkdirAll(dir, os.ModePerm)

	var scriptFile *zip.File = nil
	for _, file := range zipReader.File {
		if filepath.Ext(strings.TrimSpace(file.Name)) == ".png" {
			err := utils.CopyZipFile(file, filepath.Join(dir, filepath.Base(file.Name)))
			if err != nil {
				return "", err
			}
		} else if file.Name == "Script.json" {
			scriptFile = file
		}
	}

	scriptJSON, err := utils.ReadAllFromZip(scriptFile)
	if err != nil {
		return "", err
	}

	var rs rawScript
	json.Unmarshal([]byte(scriptJSON), &rs)

	id, err := s.repository.AddScript(ctx, name, convertToFrames(name, rs))
	if err != nil {
		return "", err
	}
	return id, nil
}

func convertToFrames(name string, rs rawScript) []editorsvc.Frame {
	frames := make([]editorsvc.Frame, len(rs.Frames))
	for i, frame := range rs.Frames {
		frames[i] = editorsvc.Frame{
			UID:         strconv.Itoa(frame.FrameNumber),
			PictureLink: filepath.Join(fmt.Sprintf("images/%s/", name), filepath.Base(frame.PictureLink)),
			Task: editorsvc.Task{
				Text: frame.Task,
			},
			Hint: editorsvc.Hint{
				Text: frame.Hint,
			},
		}
	}
	for i, frame := range rs.Frames[1:] {
		action := frame.ActionSwitch
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
	return frames
}
