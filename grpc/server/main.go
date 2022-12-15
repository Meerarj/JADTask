package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "studentsapp.com/grpc/protos"
	"studentsapp.com/grpc/server/pkg/students"
)

const (
	port = ":50052"
)

func main() {
	students.InitStudents()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterStudentServer(s, &students.StudentServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
