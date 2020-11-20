package implementation

import (
	"context"
	"io"

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
	} else if len(script) == 0 {
		return nil, editorsvc.ErrScriptNotFound
	}
	return &script[0], err
}

func (s *service) DeleteScript(ctx context.Context, id string) error {
	return s.repository.DeleteScript(ctx, id)
}

func (s *service) UpdateScript(ctx context.Context, id string, script *editorsvc.Script) (string, map[string]string, error) {
	version, err := s.repository.GetScriptVersion(ctx, id)
	if err != nil {
		return "", nil, err
	} else if version != script.Version {
		return "", nil, editorsvc.ErrVersionsDoNotMatch
	}

	script.UID = id
	script.Version = utils.RandSeq(versionLen)
	uids, err := s.repository.UpdateScript(ctx, script)
	if err != nil {
		return "", uids, err
	}
	return script.Version, uids, nil
}

func (s *service) CopyScript(ctx context.Context, script *editorsvc.Script) (string, error) {
	script.Version = utils.RandSeq(versionLen)
	return s.repository.AddScript(ctx, script)
}

func (s *service) AddBranch(ctx context.Context, script *editorsvc.Script, branch *editorsvc.Branch) (string, map[string]string, error) {
	version, err := s.repository.GetScriptVersion(ctx, script.UID)
	if err != nil {
		return "", nil, err
	} else if version != script.Version {
		return "", nil, editorsvc.ErrVersionsDoNotMatch
	}

	script.Version = utils.RandSeq(versionLen)
	uids, err := s.repository.AddBranch(ctx, script, branch)
	if err != nil {
		return "", uids, err
	}
	return script.Version, uids, nil
}

func (s *service) DeleteBranch(ctx context.Context, script *editorsvc.Script, branchToDelete *editorsvc.BranchToDelete) (string, error) {
	version, err := s.repository.GetScriptVersion(ctx, script.UID)
	if err != nil {
		return "", err
	} else if version != script.Version {
		return "", editorsvc.ErrVersionsDoNotMatch
	}

	script.Version = utils.RandSeq(versionLen)
	if err := s.repository.DeleteBranch(ctx, script, branchToDelete); err != nil {
		return "", err
	}
	return script.Version, nil
}

func (s *service) DeleteFrame(ctx context.Context, script *editorsvc.Script, id string) (string, error) {
	version, err := s.repository.GetScriptVersion(ctx, script.UID)
	if err != nil {
		return "", err
	} else if version != script.Version {
		return "", editorsvc.ErrVersionsDoNotMatch
	}

	script.Version = utils.RandSeq(versionLen)
	if err := s.repository.DeleteFrame(ctx, script, id); err != nil {
		return "", err
	}
	return script.Version, nil
}
