package main

type DBGetter interface {
	Get() string
	GetByID(id string) string
}

type dbGetter struct {
	allDbOperations
}

func NewDBGetter(d allDbOperations) DBGetter {
	return &dbGetter{allDbOperations: d}
}

func (d dbGetter) Get() string {
	return d.allDbOperations.Get()
}

func (d dbGetter) GetByID(id string) string {
	return d.allDbOperations.GetByID(id)
}
