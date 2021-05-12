package room

import (
	"cilicili-go/db/room"
	"cilicili-go/model/entity"
)

type BaseRoomServiceInterface interface {
	Create(room *entity.Room)
	FindAll() []*entity.Room
	FindById(id string) *entity.Room
	FindByLiving(living bool) []*entity.Room
	FindByStatus(status int) []*entity.Room
}

var roomDao = new(room.BaseRoomDao)

func (b *BaseRoomService) Create(room *entity.Room) {
	roomDao.Save(room)
}

func (b *BaseRoomService) FindAll() []*entity.Room {
	return roomDao.FindAll()
}

func (b *BaseRoomService) FindById(id string) *entity.Room {
	return roomDao.FindById(id)
}

func (b *BaseRoomService) FindByStatus(status int) []*entity.Room {
	return roomDao.FindByStatus(status)
}

func (b *BaseRoomService) FindByLiving(living bool) []*entity.Room {
	return roomDao.FindByLiving(living)
}

type BaseRoomService struct {
}
