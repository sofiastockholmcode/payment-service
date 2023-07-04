package main

import (
	"buf.build/gen/go/buftestorg/itemapis/bufbuild/connect-go/item/v1/itemv1connect"
	itemv1 "buf.build/gen/go/buftestorg/itemapis/protocolbuffers/go/item/v1"
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"log"
	"net/http"
)

func main() {
	HelloHandler := http.HandlerFunc(HelloServer)
	http.HandleFunc("/", HelloHandler)
	log.Println("Listen on http://localhost:8081")
	http.ListenAndServe(":8081", nil)

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	log.Println("Call to HelloServer")
	/*
			http2client := &http.Client{}
			http2client.Transport = &http2.Transport{
				AllowHTTP: true,
				DialTLS: func(netw, addr string, cfg *tls.Config) (net.Conn, error) {
					return net.Dial(netw, addr)
				}}

		userClient := userservicev1connect.NewUserServiceClient(http2client, "http://localhost:8089")

		userServiceResponse, _ := userClient.GetUser(context.Background(), connect.NewRequest(&userservicev1.GetUserRequest{
			UserId: "1",
		}))

		log.Println("Successfully Got user", userServiceResponse.Msg.GetUser().GetName())
	*/
	itemClient := itemv1connect.NewItemStoreServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	id := r.URL.Path[1:]

	item, err := itemClient.GetItem(
		context.Background(),
		connect.NewRequest(&itemv1.GetItemRequest{
			ItemId: id,
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully Got an Item", item.Msg.GetItem().Name)

	fmt.Fprintf(w, "You succuessfully bought: %s!", item.Msg.GetItem().Name)

}
