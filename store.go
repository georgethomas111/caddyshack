// This file contains the abstraction of the store entities.
package caddyshack

import (
	"github.com/bushwood/caddyshack/model"
)

type StoreObject interface {
	GetKey() string
	SetKey(string)
}

// Store represents an abstraction(interface) to any store.
// Configuration varies and let's leave that at the jurisdiction of each Store.
type Store interface {
	GetName() string
	SetName(string) error
	Init(*model.Definition) (error, Store)
	Create(StoreObject) error
	ReadOne(string) (error, StoreObject)
	UpdateOne(StoreObject) error
	DestroyOne(string) error

	//Read(query.Definition) (interface{}, error)
	//Update(interface{}) error
	//Destroy(interface{}) error
}