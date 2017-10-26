//go:generate protoc --go_out=. beeswax/augment/augmentor.proto
//go:generate protoc --go_out=. beeswax/base/eventid.proto
//go:generate protoc --go_out=. beeswax/bid/adcandidate.proto beeswax/bid/request.proto
//go:generate protoc --go_out=. beeswax/currency/currency.proto
//go:generate protoc --go_out=. beeswax/logs/ad_log.proto
//go:generate protoc --go_out=. beeswax/openrtb/extension.proto beeswax/openrtb/openrtb.proto beeswax/openrtb/openrtb_common.proto

package main

// need main function so "go get" don't complain
func main(){
}
