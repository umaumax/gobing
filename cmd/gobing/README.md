

### generate code
```
cat web.json   | gojson -pkg=gobing -name=WebResult   > web_json.go
cat image.json | gojson -pkg=gobing -name=ImageResult > image_json.go
```

