package grpc_server

import (
	"TaskManager/api/task"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	listener    net.Listener
}

func NewServer(port string, handler *Handler) (*Server, error) {
	s := &Server{}
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}
	s.listener = listener

	s.grpcServer = grpc.NewServer()
	task.RegisterTaskServiceServer(s.grpcServer, handler)

	return s, nil
}

func (s *Server) Run() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *Server) Shutdown() {
	s.grpcServer.GracefulStop()
}
