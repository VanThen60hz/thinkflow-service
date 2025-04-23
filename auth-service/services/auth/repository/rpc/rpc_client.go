package rpc

import (
	"thinkflow-service/proto/pb"

	"gorm.io/gorm"
)

type rpcClient struct {
	client pb.UserServiceClient
	db     *gorm.DB
}

func NewClient(client pb.UserServiceClient) *rpcClient {
	return &rpcClient{client: client}
}

func (c *rpcClient) GetDB() *gorm.DB {
	return c.db
}

func (c *rpcClient) SetDB(db *gorm.DB) {
	c.db = db
}
