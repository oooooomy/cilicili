package upload

import (
	"cilicili-go/db/upload"
	"cilicili-go/model/entity"
)

type BaseMusicService interface {
	Create(music *entity.Music)
	Delete(id string)
	FindAll() []*entity.Music
}

type MusicService struct{}

var musicDao = new(upload.MusicDao)

func (s *MusicService) Create(music *entity.Music) {
	musicDao.Save(music)
}

func (s *MusicService) Delete(id string) {
	musicDao.Delete(id)
}

func (s *MusicService) FindAll() []*entity.Music {
	return musicDao.FindAll()
}
