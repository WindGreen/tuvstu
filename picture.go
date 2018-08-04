package tuvstu

import (
	"os"
	"path"
	"strings"

	"github.com/globalsign/mgo/bson"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

type Picture struct {
	Path    string `json:"path" bson:"path"`
	Name    string `json:"name" bson:"name"`
	Text    string `json:"Text" bson:"text"`
	Content string `json:"content" bson:"content"`
}

func NewPicture(origin, dir string) *Picture {
	ext := path.Ext(origin)
	id, _ := uuid.NewV4()
	_, err := os.Stat(dir)
	if err != nil {
		os.MkdirAll(dir, os.ModePerm)
	}
	return &Picture{
		Name: id.String() + ext,
		Path: dir,
	}
}

func (p *Picture) GetLocation() string {
	return strings.TrimRight(p.Path, "/") + "/" + p.Name
}

func (p *Picture) Save() error {
	session := GetMgo()
	if session == nil {
		return errors.New("connection failed")
	}
	_, err := session.DB("tuvstu").C("picture").Upsert(bson.M{"name": p.Name}, p)
	return err
}

func FindPicture(name string) (*Picture, error) {
	session := GetMgo()
	if session == nil {
		return nil, errors.New("connection failed")
	}
	picture := new(Picture)
	err := session.DB("tuvstu").C("picture").Find(bson.M{"name": name}).One(picture)
	if err != nil {
		return nil, err
	}
	return picture, nil
}
