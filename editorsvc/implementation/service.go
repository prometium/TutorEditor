package implementation

import (
	"context"
	"io"
	"time"

	"github.com/prometium/tutoreditor/editorsvc"
	"github.com/prometium/tutoreditor/editorsvc/utils"
)

type service struct {
	repository editorsvc.Repository
}

const versionLen = 10

// NewService makes a new Service
func NewService(rep editorsvc.Repository) editorsvc.Service {
	return &service{
		repository: rep,
	}
}

func (s *service) AddRawScript(ctx context.Context, name string, fileReader io.ReadCloser) (string, error) {
	if fileReader == nil {
		return "", editorsvc.ErrFileNotAttached
	}

	defer fileReader.Close()

	var rs rawScript
	if err := rs.init(fileReader); err != nil {
		return "", err
	}

	linksMap, err := rs.saveImages(ctx, "assets/images/")
	if err != nil {
		return "", err
	}

	script, err := rs.createScript(name, linksMap)
	if err != nil {
		return "", err
	}

	script.Version = utils.RandSeq(versionLen)
	script.ModificationDate = time.Now().Format("2006.01.02 15:04:05")
	id, err := s.repository.AddScript(ctx, script)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *service) GetScriptsList(ctx context.Context) ([]editorsvc.Script, error) {
	return s.repository.GetScriptsList(ctx)
}

func (s *service) GetScript(ctx context.Context, id string) (*editorsvc.Script, error) {
	script, err := s.repository.GetScript(ctx, id)
	if err != nil {
		return nil, err
	} else if script == nil {
		return nil, editorsvc.ErrScriptNotFound
	}
	return script, err
}

func (s *service) DeleteScript(ctx context.Context, id string) error {
	return s.repository.DeleteScript(ctx, id)
}

func (s *service) UpdateScript(
	ctx context.Context, script *editorsvc.Script, frameIdsToDel []string, actionIdsToDel []string,
) (map[string]string, error) {
	script.Version = utils.RandSeq(versionLen)
	script.ModificationDate = time.Now().Format("2006.01.02 15:04:05")
	uids, err := s.repository.UpdateScript(ctx, script, frameIdsToDel, actionIdsToDel)
	if err != nil {
		return uids, err
	}
	return uids, nil
}

func (s *service) CopyScript(ctx context.Context, script *editorsvc.Script) (string, error) {
	script.Version = utils.RandSeq(versionLen)
	return s.repository.AddScript(ctx, script)
}
