package dao

import (
	"daimon/api"
	"daimon/models"
	"daimon/pkg/mysql"

	"gorm.io/gorm"
)

type ProjectDao struct {
	*gorm.DB
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{mysql.GetDB()}
}

func (d *ProjectDao) AddProject(p *models.Project) error {
	return d.Create(p).Error
}

func (d *ProjectDao) DelProject(id uint) {
	d.Where("id = ?", id).Delete(&models.Project{})
}

func (d *ProjectDao) UpdateProject(p *models.Project) error {
	return d.Model(&models.Project{}).Where("id = ?", p.ID).Omit("id, deleted_at").Updates(p).Error
}

func (d *ProjectDao) GetProjectById(id uint) (p *models.Project, err error) {
	p = new(models.Project)
	err = d.Model(&models.Project{}).Where("id = ?", id).First(p).Error
	return
}

func (d *ProjectDao) ListAllProjects() (ps []*api.Project) {
	d.Model(&models.Project{}).Order("created_at DESC").Find(&ps)
	return
}

func (d *ProjectDao) GetProjectDetail(id uint) (p *api.Project, err error) {
	p = new(api.Project)
	d.Model(&models.Project{}).Where("id = ?", id).First(p)
	var tgs []*api.TaskGroup
	d.Model(&models.TaskGroup{}).Where("project_id = ?", id).Order("serial ASC").Find(&tgs)
	var ts []*api.Task
	d.Model(&models.Task{}).Where("project_id = ?", id).Order("task_group_id ASC, serial ASC").Find(&ts)

	idx := 0
	for _, task := range ts {

		for {
			if tgs[idx%len(tgs)].Id == task.TaskGroupId {
				break
			} else {
				idx += 1
			}
		}
		tgs[idx%len(tgs)].Tasks = append(tgs[idx%len(tgs)].Tasks, task)
	}

	p.TaskGroups = append(p.TaskGroups, tgs...)

	return
}
