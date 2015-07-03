package db

import (
	"gopkg.in/mgo.v2"
)

type Client struct {
	*mgo.Collection
}

func New(address, dbName, collection string) (*Client, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	c := session.DB(dbName).C(collection)
	return &Client{c}, nil
}

func (self *Client) GetUsersByApp(appName string) (users []string) {

}
