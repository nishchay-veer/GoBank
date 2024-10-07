package gapi

import (
	"log"

	db "github.com/nishchay-veer/simplebank/db/sqlc"
	"github.com/nishchay-veer/simplebank/pb"
	"github.com/nishchay-veer/simplebank/token"
	"github.com/nishchay-veer/simplebank/util"
	"github.com/nishchay-veer/simplebank/worker"
)

// This server is going to serve gRPC  requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPASETOMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker")

	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
