package convert

import (
	"fmt"
	pbg "github.com/grzegab/gRPC/proto_examples"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func convertToFile(n string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = os.WriteFile(n, out, 0644); err != nil {
		log.Fatal(err)
		return
	}

	log.Println("File created:", n)
}

func readFromFile(n string, pb proto.Message) {
	in, err := os.ReadFile(n)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatal(err)
		return
	}
}

func doFile() {
	path := "simple.bin"
	msg := &pbg.Example{}
	convertToFile(path, msg)

	readFromFile(path, msg)
	fmt.Println(msg)
}
