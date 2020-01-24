package app

import (
	"github.com/webitel/engine/auth_manager"
	"github.com/webitel/engine/model"
)

func (a *App) CreateList(list *model.List) (*model.List, *model.AppError) {
	return a.Store.List().Create(list)
}

func (a *App) ListCheckAccess(domainId, id int64, groups []int, access auth_manager.PermissionAccess) (bool, *model.AppError) {
	return a.Store.List().CheckAccess(domainId, id, groups, access)
}

func (a *App) GetListPage(domainId int64, page, perPage int) ([]*model.List, *model.AppError) {
	return a.Store.List().GetAllPage(domainId, page*perPage, perPage)
}

func (a *App) GetListPageByGroups(domainId int64, groups []int, page, perPage int) ([]*model.List, *model.AppError) {
	return a.Store.List().GetAllPageByGroups(domainId, groups, page*perPage, perPage)
}

func (a *App) GetListById(domainId, id int64) (*model.List, *model.AppError) {
	return a.Store.List().Get(domainId, id)
}

func (a *App) UpdateList(list *model.List) (*model.List, *model.AppError) {
	oldList, err := a.GetListById(list.DomainId, list.Id)
	if err != nil {
		return nil, err
	}

	oldList.Description = list.Description
	oldList.Name = list.Name
	oldList.UpdatedAt = list.UpdatedAt
	oldList.UpdatedBy.Id = list.UpdatedBy.Id

	oldList, err = a.Store.List().Update(oldList)
	if err != nil {
		return nil, err
	}

	return oldList, nil
}

func (a *App) RemoveList(domainId, id int64) (*model.List, *model.AppError) {
	list, err := a.Store.List().Get(domainId, id)

	if err != nil {
		return nil, err
	}

	err = a.Store.List().Delete(domainId, id)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (a *App) CreateListCommunication(comm *model.ListCommunication) (*model.ListCommunication, *model.AppError) {
	return a.Store.List().CreateCommunication(comm)
}

func (a *App) GetListCommunicationPage(domainId, listId int64, page, perPage int) ([]*model.ListCommunication, *model.AppError) {
	return a.Store.List().GetAllPageCommunication(domainId, listId, page*perPage, perPage)
}

func (a *App) GetListCommunicationById(domainId, listId, id int64) (*model.ListCommunication, *model.AppError) {
	return a.Store.List().GetCommunication(domainId, listId, id)
}

func (a *App) UpdateListCommunication(domainId int64, communication *model.ListCommunication) (*model.ListCommunication, *model.AppError) {
	oldComm, err := a.GetListCommunicationById(domainId, communication.ListId, communication.Id)
	if err != nil {
		return nil, err
	}

	oldComm.Description = communication.Description
	oldComm.Number = communication.Number

	oldComm, err = a.Store.List().UpdateCommunication(domainId, oldComm)
	if err != nil {
		return nil, err
	}

	return oldComm, nil
}

func (a *App) RemoveListCommunication(domainId, listId, id int64) (*model.ListCommunication, *model.AppError) {
	communication, err := a.Store.List().GetCommunication(domainId, listId, id)

	if err != nil {
		return nil, err
	}

	err = a.Store.List().DeleteCommunication(domainId, listId, id)
	if err != nil {
		return nil, err
	}
	return communication, nil
}
