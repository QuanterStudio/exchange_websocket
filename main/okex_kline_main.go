package main

import . "exchange_websocket/okex_websocket"

func main() {
	okex := OkexWebsocketInit()
	okex.OkexKlineWebsocket()
	for {
		okex.WsConnect()
		go func() {
			okex.Ping()
		}()
		okex.Subscribe()
		okex.ReadMessage()

	}

}
