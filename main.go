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

/*func handle(w http.ResponseWriter, r *http.Request) {
	augRequest := augment.AugmentorRequest{}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	if err := proto.Unmarshal(data, &augRequest); err != nil {
		fmt.Println(err)
	}

	user := augRequest.GetBidRequest().GetUser()
	fmt.Println("user id:" + user.GetId())

	// unknow user
	if user.Id == nil {
		w.WriteHeader(http.StatusNoContent)
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
		w.WriteHeader(http.StatusNoContent)
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
	w.Write(respByte)
}*/
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	augRequest := augment.AugmentorRequest{}
	if err := proto.Unmarshal(ctx.PostBody(), &augRequest); err != nil {
		fmt.Println(err)
	}

	user := augRequest.GetBidRequest().GetUser()
	fmt.Println("user id:" + user.GetId())

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
	//w.Write(respByte)
}

func main() {
	//http.HandleFunc("/", handle)
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
	//http.ListenAndServe(":3000", nil)
	fasthttp.ListenAndServe(":3000", fastHTTPHandler)
}
