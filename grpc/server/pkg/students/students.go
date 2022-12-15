package students

import (
	"context"
	"log"
	"sync"

	pb "studentsapp.com/grpc/protos"
)

var students []*pb.StudentInfo

type StudentServer struct {
	pb.UnimplementedStudentServer
	m sync.Mutex
}

func InitStudents() {
	student1 := &pb.StudentInfo{Id: "1", Rollno: 1, Name: "Meera",
		Class: 11, Section: "A", Hobbies: []string{"Painting", "Reading"}}
	student2 := &pb.StudentInfo{Id: "2", Rollno: 2, Name: "Devi",
		Class: 12, Section: "B", Hobbies: []string{"Singing", "Dancing"}}

	students = append(students, student1)
	students = append(students, student2)
}
func (s *StudentServer) GetStudents(in *pb.Empty,
	stream pb.Student_GetStudentsServer) error {
	log.Printf("Received: %v", in)
	s.m.Lock()
	for _, student := range students {
		if err := stream.Send(student); err != nil {
			return err

		}
	}
	s.m.Unlock()

	return nil
}
func (s *StudentServer) GetStudentBySection(ctx context.Context, in *pb.Section) (*pb.StudentInfo, error) {
	log.Printf("Received: %v", in)

	res := &pb.StudentInfo{}
	s.m.Lock()
	for _, student := range students {
		if student.GetSection() == in.GetValue() {
			res = student
			break
		} else {
			log.Printf("No student exist with this section")
		}
	}
	s.m.Unlock()
	return res, nil
}

func (s *StudentServer) GetStudentByClass(ctx context.Context, in *pb.Class) (*pb.StudentInfo, error) {
	log.Printf("Received: %v", in)

	res := &pb.StudentInfo{}
	s.m.Lock()

	for _, student := range students {
		if student.GetClass() == in.GetValue() {
			res = student
			students = append(students, res)
			// break
		} else {
			log.Printf("No student exist in this class")
		}
	}
	s.m.Unlock()
	return res, nil
}

func (s *StudentServer) CreateStudent(ctx context.Context,
	in *pb.StudentInfo) (*pb.Id, error) {
	log.Printf("Received: %v", in)
	s.m.Lock()
	defer s.m.Unlock()
	res := pb.Id{}
	// res.Value = strconv.Itoa(rand.Intn(100000000))
	res.Value = in.Id

	in.Id = res.GetValue()
	students = append(students, in)
	return &res, nil
}
func (s *StudentServer) UpdateStudent(ctx context.Context,
	in *pb.StudentInfo) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	s.m.Lock()
	defer s.m.Unlock()
	res := pb.Status{}
	for index, student := range students {
		if student.GetId() == in.GetId() {
			students = append(students[:index], students[index+1:]...)
			in.Id = student.GetId()
			students = append(students, in)
			res.Value = 1
			break
		}
	}

	return &res, nil
}

func (s *StudentServer) DeleteStudent(ctx context.Context,
	in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	s.m.Lock()
	defer s.m.Unlock()
	res := pb.Status{}
	for index, student := range students {
		if student.GetId() == in.GetValue() {
			students = append(students[:index], students[index+1:]...)
			res.Value = 1
			break
		} else {
			log.Printf("No records in this id")
		}
	}

	return &res, nil
}
