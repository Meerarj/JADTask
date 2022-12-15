package main

import (
	"fmt"

	"log"

	"google.golang.org/grpc"
	"studentsapp.com/grpc/client/pkg/students"
	pb "studentsapp.com/grpc/protos"
)

const (
	address = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewStudentClient(conn)
	fmt.Println(client)

	students.RunGetStudents(client)
	// students.RunGetStudentBySection(client, "B")
	// students.RunGetStudentByClass(client, 11)
	// students.RunCreateStudent(client, "5", 6, "Sachi", 11, "B", []string{"Skating", "Gardening"})
	// students.RunCreateStudent(client, "3", 7, "Ram", 11, "B", []string{"Skating", "Gardening"})

	// students.RunUpdateStudent(client, "1", 1, "Renju", 12, "C", []string{"Skating", "Singing"})

	// students.RunDeleteStudent(client, "2")
	// }
}
