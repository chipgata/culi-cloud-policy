package main

import (
	"fmt"
	"github.com/pborman/uuid"
	"strings"
)

func main()  {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
}