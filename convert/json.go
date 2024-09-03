package convert

import (
	pb "github.com/grzegab/gRPC/greet/proto"
	"google.golang.org/protobuf/proto"
	"log"
	"reflect"
)

func toJSON(pb proto.Message) string {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return string(out)
}

func fromJSON(json string, pb proto.Message) {
	if err := proto.Unmarshal([]byte(json), pb); err != nil {
		log.Fatal(err)
	}

}

func doToJSON(pb proto.Message) string {
	s := toJSON(pb)
	log.Println(s)

	return s
}

func doFromJSON(json string, t reflect.Type) proto.Message {
	msg := reflect.New(t).Interface().(proto.Message)

	fromJSON(json, msg)

	return msg
}

func doIt() {
	message := &pb.GreetRequest{}
	jsonString := doToJSON(message)

	newMessage := doFromJSON(jsonString, reflect.TypeOf(message))
	log.Println(newMessage)
}
