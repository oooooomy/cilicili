package upload

import (
	"cilicili-go/model/entity"
	"github.com/go-basic/uuid"
)

type BaseMusicDao interface {
	Save(music *entity.Music)
	Delete(id string)
	FindAll() []*entity.Music
}

type MusicDao struct{}

func (b *MusicDao) Save(music *entity.Music) {
	music.Id = uuid.New()
	conn.Table("t_music").Save(music)
}

func (b *MusicDao) Delete(id string) {
	conn.Table("t_music").Delete(&entity.Music{}, "id = ?", id)
}

func (b *MusicDao) FindAll() (musics []*entity.Music) {
	conn.Table("t_music").Find(&musics)
	return
}
