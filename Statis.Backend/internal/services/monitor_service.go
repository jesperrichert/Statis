package services

import (
	"net/http"
	"strings"

	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/schedule"
	"jespersen.zip.statis/internal/types/monitors"
	"jespersen.zip.statis/internal/util"
)

type MonitorService struct {
	MonitorScheduler *schedule.MonitorScheduler
	DB               *gorm.DB
}

func NewMonitorService(db *gorm.DB, scheduler *schedule.MonitorScheduler) *MonitorService {
	return &MonitorService{
		DB:               db,
		MonitorScheduler: scheduler,
	}
}

func (e *MonitorService) Get(identifier string) (dto.MonitorDto[any], bool) {
	var monitor model.MonitorModel
	e.DB.Where("identifier = ?", identifier).First(&monitor)
	if util.CheckNull(monitor.Identifier) {
		return dto.MonitorDto[any]{}, false
	}

	e.DB.Preload("Events").Where("identifier = ?", identifier).First(&monitor)

	var affectedBy []dto.MonitorEffect
	for k, v := range monitor.AffectedByExternal {
		affectedBy = append(affectedBy, dto.MonitorEffect{
			Name: k,
			Link: v,
		})
	}

	var events []string
	for i := range monitor.Events {
		event := monitor.Events[i]
		events = append(events, event.Identifier)
	}

	return dto.MonitorDto[any]{
		Identifier:   monitor.Identifier,
		IsParent:     monitor.IsParent,
		Group:        monitor.Group,
		Name:         monitor.Name,
		Description:  monitor.Description,
		Type:         monitor.Type,
		ScheduleTime: monitor.ScheduleTime,
		TypeData:     monitor.TypeData,
		State:        monitor.State,
		Events:       events,
		AffectedBy:   affectedBy,
		ExternalUrl:  monitor.ExternalUrl,
	}, true
}

func (e *MonitorService) Delete(identifier string) bool {
	var monitor model.MonitorModel
	e.DB.Preload("Events").Where("identifier = ?", identifier).First(&monitor)
	if util.CheckNull(monitor.Identifier) {
		return false
	}

	e.DB.Delete(&monitor)
	return true
}

func (e *MonitorService) Validate(dto *dto.MonitorDto[any]) (int, string) {

	if util.CheckNull(dto.Identifier) ||
		util.CheckNull(dto.IsParent) ||
		util.CheckNull(dto.Name) ||
		util.CheckNull(dto.ScheduleTime) ||
		util.CheckNull(dto.Type) ||
		util.CheckNull(dto.TypeData) {
		return http.StatusBadRequest, "dto not valid"
	}

	if !monitors.ToType(dto.Type).IsTypeDtoValid(dto) ||
		dto.TypeData == nil ||
		util.IsEmpty(dto.TypeData) {
		return http.StatusBadRequest, "invalid monitor type"
	}

	var monitorModel model.MonitorModel
	e.DB.Where(&model.MonitorModel{}).Where("identifier = ?", dto.Identifier).First(&monitorModel)
	if !util.CheckNull(monitorModel.Identifier) {
		return http.StatusConflict, "monitor already exists!"
	}

	return http.StatusOK, ""
}

func (e *MonitorService) Add(dto *dto.MonitorDto[any]) int {
	var effects = map[string]string{}
	for i := range dto.AffectedBy {
		effects[dto.AffectedBy[i].Name] = dto.AffectedBy[i].Link
	}

	var monitorModel model.MonitorModel
	e.DB.Create(&model.MonitorModel{
		Identifier:         dto.Identifier,
		IsParent:           dto.IsParent,
		Group:              dto.Group,
		Name:               dto.Name,
		Description:        dto.Description,
		ScheduleTime:       dto.ScheduleTime,
		Type:               strings.ToUpper(dto.Type),
		ExternalUrl:        dto.ExternalUrl,
		TypeData:           dto.TypeData,
		State:              monitors.Offline.ToString(),
		AffectedByExternal: effects,
	}).Where("identifier = ?", dto.Identifier).First(&monitorModel)

	monitor := model.Monitor[any](monitorModel)
	e.MonitorScheduler.ScheduleMonitor(&monitor)

	return http.StatusOK
}

func (e *MonitorService) Update(identifier string, dto *dto.MonitorDto[any]) int {

	var monitorModel model.MonitorModel
	e.DB.Where("identifier = ?", identifier).First(&monitorModel)
	if util.CheckNull(monitorModel.Identifier) {
		return http.StatusNotFound
	}

	if !util.CheckNull(dto.Type) && util.CheckNull(dto.TypeData) {
		monitorModel.Type = strings.ToUpper(dto.Type)
		monitorModel.TypeData = dto.TypeData
	}
	if !util.CheckNull(dto.Name) {
		monitorModel.Name = dto.Name
	}
	if !util.CheckNull(dto.Description) {
		monitorModel.Description = dto.Description
	}
	if !util.CheckNull(dto.AffectedBy) {
		effects := map[string]string{}
		for i := range dto.AffectedBy {
			effect := dto.AffectedBy[i]
			effects[effect.Name] = effect.Link
		}
		monitorModel.AffectedByExternal = effects
	}
	if !util.CheckNull(dto.ScheduleTime) {
		monitorModel.ScheduleTime = dto.ScheduleTime
	}
	if !util.CheckNull(dto.ExternalUrl) {
		monitorModel.ExternalUrl = dto.ExternalUrl
	}
	if !util.CheckNull(dto.IsParent) {
		monitorModel.IsParent = dto.IsParent
		if dto.IsParent {
			if !util.CheckNull(dto.Group) {
				monitorModel.Group = dto.Group
			}
		} else {
			monitorModel.Group = ""
		}
	}

	key := e.MonitorScheduler.Key((*model.Monitor[any])(&monitorModel))
	if e.MonitorScheduler.Jobs[key] != nil {
		_ = e.MonitorScheduler.Scheduler.RemoveJob(e.MonitorScheduler.Jobs[key].ID())
	}
	monitor := model.Monitor[any](monitorModel)
	e.MonitorScheduler.ScheduleMonitor(&monitor)

	e.DB.Where("id = ?", monitorModel.ID).Updates(&monitorModel)

	return http.StatusOK
}
