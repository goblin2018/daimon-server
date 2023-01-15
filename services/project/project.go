package project

import (
	"daimon/api"
	"daimon/dao"
	"daimon/pkg/ctx"
)

type ProjectService struct {
	dao *dao.ProjectDao
}

func NewProjectService() *ProjectService {
	return &ProjectService{dao.NewProjectDao()}
}

func (s *ProjectService) AddProject(c *ctx.Context, req *api.Project) (res *api.Project, err error) {
	res = new(api.Project)
	// err = s.dao.AddProject(&models.Project{
	// 	Name:    req.Name,
	// 	Desc:    req.Desc,
	// 	StartAt: req.StartAt,
	// 	EndAt:   req.EndAt,
	// })

	res.Id = 2
	res.Name = "test"
	res.Desc = "this is a test message reply from post method"

	return
}

func (s *ProjectService) DelProject(c *ctx.Context, req *api.Project) (err error) {
	s.dao.DelProject(req.Id)
	return nil
}

func (s *ProjectService) ListProjects(c *ctx.Context) (ps []*api.Project) {
	ps = s.dao.ListAllProjects()
	return
}

func (s *ProjectService) GetProjectInfo(c *ctx.Context, req *api.Project) (resp *api.Project, err error) {
	// resp, err = s.dao.GetProjectDetail(req.Id)
	resp = &api.Project{
		Id:   2,
		Name: "test info",
		Desc: "this is a test message reply from get method",
	}
	return
}
