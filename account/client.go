package account

import (
	"context"

	"github.com/rajan-marasini/ecom-microservice/account/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.NewClient(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := pb.NewAccountServiceClient(conn)

	return &Client{
		conn,
		c,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostAccount(ctx context.Context, name string) (*Account, error) {
	r, err := c.service.PostAccount(
		ctx,
		&pb.PostAccountRequest{
			Name: name,
		})

	if err != nil {
		return nil, err
	}

	return &Account{
		ID:   r.Account.Id,
		Name: r.Account.Name,
	}, nil
}

func (c *Client) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	r, err := c.service.GetAccountByID(
		ctx,
		&pb.GetAccountByIDRequest{
			Id: id,
		})
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:   r.Account.Id,
		Name: r.Account.Name,
	}, nil
}

func (c *Client) GetAccounts(ctx context.Context, take uint64, skip uint64) ([]Account, error) {
	r, err := c.service.GetAccounts(
		ctx,
		&pb.GetAccountsRequest{
			Skip: skip,
			Take: take,
		})
	if err != nil {
		return nil, err
	}

	accounts := []Account{}

	for _, a := range r.Accounts {
		accounts = append(accounts, Account{
			ID:   a.Id,
			Name: a.Name,
		})
	}

	return accounts, nil

}
