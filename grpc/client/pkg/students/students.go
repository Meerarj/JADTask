package students

import (
	"context"
	"io"
	"log"
	"time"

	pb "studentsapp.com/grpc/protos"
)

func RunGetStudents(client pb.StudentClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Empty{}
	stream, err := client.GetStudents(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetStudents(_) = _, %v", client, err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetStudents(_) = _, %v", client, err)
		}
		log.Printf("StudentInfo: %v", row)
	}
}

func RunGetStudentBySection(client pb.StudentClient, studentsection string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := &pb.Section{Value: studentsection}

	res, err := client.GetStudentBySection(ctx, req)

	if err != nil {
		log.Fatalf("%v.GetStudentBySection(_) = _, %v", client, err)
	}
	log.Printf("StudentDetails: %v", res)
}

func RunGetStudentByClass(client pb.StudentClient, studentclass int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := &pb.Class{Value: studentclass}
	res, err := client.GetStudentByClass(ctx, req)

	if err != nil {
		log.Fatalf("%v.GetStudentByClass(_) = _, %v", client, err)

	}
	// for{
	// 	row,err:=res.Recv()
	// 	if err==io.EOF{
	// 		break
	// 	}
	// }
	log.Printf("StudentDetails: %v", res)
}

func RunCreateStudent(client pb.StudentClient, id string, rollno int32, name string,
	class int32, section string, hobbies []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.StudentInfo{Id: id, Rollno: rollno, Name: name, Class: class, Section: section, Hobbies: hobbies}
	res, err := client.CreateStudent(ctx, req)
	if err != nil {
		log.Fatalf("%v.CreateStudent(_) = _, %v", client, err)
	}
	if res.GetValue() != "" {
		log.Printf("CreateStudent id: %v", res)
	} else {
		log.Printf("CreateStudent Failed")
	}

}

func RunUpdateStudent(client pb.StudentClient, id string, rollno int32,
	name string, class int32, section string, hobbies []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.StudentInfo{Id: id, Rollno: rollno, Name: name,
		Class: class, Section: section, Hobbies: hobbies}
	res, err := client.UpdateStudent(ctx, req)
	if err != nil {
		log.Fatalf("%v.UpdateStudent(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("UpdateStudent Success")
	} else {

		log.Printf("UpdateStudent Failed %v", res.GetValue())
	}
}

func RunDeleteStudent(client pb.StudentClient, id string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: id}
	res, err := client.DeleteStudent(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteStudent(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("Deletestudent Success")
	} else {
		log.Printf("DeleteStudent Failed")
	}
}
