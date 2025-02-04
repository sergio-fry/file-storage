# File Storage

Service to store large files with HTTP API.

Upload file:

```bash
curl -X POST -H "Content-Type:multipart/form-data" -F "file=@file.mp4" http://0.0.0.0:8080/upload
```

Download file:

```bash
curl http://0.0.0.0:8080/files\?name\=file.mp4
```
