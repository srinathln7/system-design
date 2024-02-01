# Web-Socket


## LiveFeed

To access feed, open your web browser and navigate to the console (F12)

```
feedSocket = new WebSocket("ws://{{.Host}}/feed")
feedSocket.onmessage = (event) => {console.log("received from feed server", event.data)}
```


## Chat application

Refer this [video](https://www.youtube.com/watch?v=y036l6pvVEs) 