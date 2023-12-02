package main

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/lixvyang/rum-sdk-go/sdk/base"
	"github.com/lixvyang/rum-sdk-go/sdk/groups"
	"github.com/lixvyang/rum-sdk-go/sdk/model"
	"github.com/lixvyang/rum-sdk-go/service/fullnode"
)

func main() {
	cli := fullnode.New(&base.Configuration{
		BasePath: "http://127.0.0.1:8000",
	})

	// _, _, err := cli.AppsApi.AppApiV1GroupGroupIdContentGet(context.Background(), "cd596d77-b789-459f-8642-a469279d11a6", "cd596d77-b789-459f-8642-a469279d11a6", &apps.AppsApiAppApiV1GroupGroupIdContentGetOpts{
	// 	Num:      optional.NewInt64(10),
	// 	Reverse:  optional.NewBool(true),
	// 	StartTrx: optional.NewString("bfa6475d-a56b-49e7-a751-d459e981de0c"),
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// resp, _, err := cli.GroupsApi.ApiV1GroupPost(context.Background(), &groups.GroupsApiApiV1GroupPostOpts{
	// 	HandlersCreateGroupParam: optional.NewInterface(model.HandlersCreateGroupParam{
	// 		AppKey:          "123456",
	// 		ConsensusType:   "poa",
	// 		EncryptionType:  "private",
	// 		GroupName:       "test",
	// 		IncludeChainUrl: true,
	// 	}),
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v", resp)

	// resp, _, err := cli.UserApi.ApiV1GroupAnnouncePost(context.Background(), &user.UserApiApiV1GroupAnnouncePostOpts{
	// 	HandlersAnnounceParam: optional.NewInterface(model.HandlersAnnounceParam{
	// 		Action:  "add",
	// 		GroupId: "70851e7c-3ad0-44e1-91b6-85f074c2c768",
	// 		Memo:    "invatation code: 213121",
	// 		Type:    "user",
	// 	}),
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v", resp)

	resp, _, err := cli.GroupsApi.ApiV1GroupLeavePost(context.Background(), &groups.GroupsApiApiV1GroupLeavePostOpts{
		HandlersLeaveGroupParam: optional.NewInterface(model.HandlersLeaveGroupParam{
			GroupId: "70851e7c-3ad0-44e1-91b6-85f074c2c768",
		}),
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", resp)

	respp, _, err := cli.GroupsApi.ApiV1GroupClearPost(context.Background(), &groups.GroupsApiApiV1GroupClearPostOpts{
		HandlersClearGroupDataParam: optional.NewInterface(model.HandlersClearGroupDataParam{
			GroupId: "70851e7c-3ad0-44e1-91b6-85f074c2c768",
		}),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", respp)

	// respp, _, err := cli.NodeApi.ApiV1NetworkGet(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", respp)
}
