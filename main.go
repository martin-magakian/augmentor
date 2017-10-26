package main

import (
	"fmt"
	"log"
	"os"

	"beeswax/augment"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
	"github.com/valyala/fasthttp"
)

var c redis.Conn

func stringPtr(s string) *string {
	return &s
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	augRequest := augment.AugmentorRequest{}
	if err := proto.Unmarshal(ctx.PostBody(), &augRequest); err != nil {
		fmt.Println(err)
	}

	user := augRequest.GetBidRequest().GetUser()
	fmt.Println("user id :" + user.GetId())

	// unknow user
	if user.Id == nil {
		ctx.SetStatusCode(fasthttp.StatusNoContent)
		return
	}
	fmt.Println(user.GetId())
	rawUserCounter, err := c.Do("INCR", "userId-"+user.GetId())
	if err != nil {
		log.Fatalln(err)
	}
	userCounter := rawUserCounter.(int64)
	fmt.Println(userCounter)
	if userCounter == 1 { // user is unknow
		ctx.SetStatusCode(fasthttp.StatusNoContent)
		return
	}

	// user is know add Augment is segment
	response := augment.AugmentorResponse{}
	response.Segments = []*augment.AugmentorResponse_Segment{
		&augment.AugmentorResponse_Segment{
			Id:    stringPtr("intentioniste"),
			Value: stringPtr("Audi A3"),
		},
		&augment.AugmentorResponse_Segment{
			Id:    stringPtr("Intentioniste"),
			Value: stringPtr("Voiture"),
		},
	}
	respByte, err := proto.Marshal(&response)
	if err != nil {
		log.Fatalln(err)
	}
	ctx.SetBody(respByte)
}

func main() {
	connection, err := redis.Dial("tcp", "redis_db:6379")
	c = connection
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	resp, err := c.Do("PING")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	fasthttp.ListenAndServe(":3000", fastHTTPHandler)
}
