package main

import (
	"context"
	"fmt"

	"github.com/lixvyang/rum-sdk-go/sdk/base"
	"github.com/lixvyang/rum-sdk-go/service/fullnode"
)

func main() {
	cli := fullnode.New(&base.Configuration{
		BasePath: "http://127.0.0.1:8000",
	})
	resp, _, err := cli.GroupsApi.ApiV1GroupGroupIdGet(context.Background(), "c7e9b833-e60a-4901-9abc-c8ec5b15a7e6")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)

	respp, _, err := cli.NodeApi.ApiV1NetworkGet(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", respp)

}
