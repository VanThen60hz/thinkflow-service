package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
)

type rpcUserClient struct {
	client pb.UserServiceClient
}

func NewUserClient(client pb.UserServiceClient) *rpcUserClient {
	return &rpcUserClient{client: client}
}

func (c *rpcUserClient) GetUsersByIds(ctx context.Context, ids []int) ([]core.SimpleUser, error) {
	userIds := make([]int32, len(ids))

	for i := range ids {
		userIds[i] = int32(ids[i])
	}

	resp, err := c.client.GetUsersByIds(ctx, &pb.GetUsersByIdsReq{Ids: userIds})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	users := make([]core.SimpleUser, len(resp.Users))

	for i := range users {
		respUser := resp.Users[i]
		users[i] = core.NewSimpleUser(int(respUser.Id), respUser.Email, respUser.FirstName, respUser.LastName, nil)
	}

	return users, nil
}

func (c *rpcUserClient) GetUserById(ctx context.Context, id int) (*core.SimpleUser, error) {
	resp, err := c.client.GetUserById(ctx, &pb.GetUserByIdReq{Id: int32(id)})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user := core.NewSimpleUser(int(resp.User.Id), resp.User.Email, resp.User.FirstName, resp.User.LastName, nil)

	return &user, nil
}
