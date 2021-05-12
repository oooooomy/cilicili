package room

import (
	"cilicili-go/model/entity"
)

type BaseRoomDaoInterface interface {
	Save(room *entity.Room)
	FindAll() []*entity.Room
	FindById(id string) *entity.Room
	FindByStatus(status int) []*entity.Room
	FindByLiving(living bool) []*entity.Room
}

func (b *BaseRoomDao) Save(room *entity.Room) {
	conn.Table("t_room").Save(room)
}

func (b *BaseRoomDao) FindAll() (rooms []*entity.Room) {
	conn.Table("t_room").Find(&rooms)
	return
}

func (b *BaseRoomDao) FindById(id string) *entity.Room {
	room := new(entity.Room)
	conn.Table("t_room").Find(room, "id = ?", id)
	return room
}

func (b *BaseRoomDao) FindByStatus(status int) (rooms []*entity.Room) {
	conn.Table("t_room").Find(&rooms, "status = ?", status)
	return
}

func (b *BaseRoomDao) FindByLiving(living bool) (rooms []*entity.Room) {
	conn.Table("t_room").Find(&rooms, "is_living = ?", living)
	return
}

type BaseRoomDao struct{}
