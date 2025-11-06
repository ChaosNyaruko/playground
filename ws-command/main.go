// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"encoding/hex"
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	// send command
	go func() {
		t := time.NewTicker(5 * time.Second)
		for range t.C {
			cmd := Command{
				Command:      "start_task",
				TaskID:       uuid.New().String(),
				TaskType:     "one-shot",
				ConcreteType: "interface_test",
				Params:       map[string]interface{}{"biz": "vod", "method": "GetPeers"},
				Interval:     0,
			}
			err = c.WriteJSON(cmd)
			log.Println("write ", cmd)
			if err != nil {
				log.Println("err: write:", err)
				break
			}
		}
	}()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("xx read:", err)
			break
		}
		if mt == 2 {
			s := []byte(hex.EncodeToString(message))
			log.Printf("recv: %d, %s", mt, s)
		} else {
			log.Printf("recv: %d, %s", mt, message)
		}
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))

type Command struct {
	Command      string                 `json:"command"`
	TaskID       string                 `json:"task_id,omitempty"`
	TaskType     string                 `json:"task_type,omitempty"`
	ConcreteType string                 `json:"concrete_type,omitempty"`
	Params       map[string]interface{} `json:"params,omitempty"`
	Interval     int                    `json:"interval,omitempty"`
}
