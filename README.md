# top-menu-items
Reads the log file and tell the top 3 items consumed.

##API docs
1. 
Request 

## to build the program
Directly run
``` go run main.go ```

Build and run
```go build -o main && ./main```

## API docs
Default port is 8080
1. Ping Request 
```
curl --location 'localhost:8080/ping'
```

```
Response
{
    "message": "pong"
}
```


2. Upload menu logs
```
Request

curl --location 'localhost:8080/upload-menu-log' \
--header 'Content-Type: application/json' \
--data '{
    "menu_logs": {
        "menu_logs": [
            {
                "eater_id": 100,
                "foodmenu_id": 1
            },
            {
                "eater_id": 101,
                "foodmenu_id": 1
            },
            {
                "eater_id": 102,
                "foodmenu_id": 1
            },
            {
                "eater_id": 103,
                "foodmenu_id": 1
            },
            {
                "eater_id": 104,
                "foodmenu_id": 1
            }
        ]
    }
}'
```

```
Response:-

{
    "StatusCode": 200,
    "error": "",
    "data": {
        "top_food_item_ids": [
            1
        ]
    }
}
```
