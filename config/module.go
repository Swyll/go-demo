package config

type IniParser interface {
	Load(path string, conf interface{}) error
	Save(path string, conf interface{}) error
}
