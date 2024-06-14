# worker api

## sample sequence

```
if no worker conf{
    worker_id = POST /worker
}else{
    worker_id = Read conf
}

loop {
    job_id = POST /worker/$worker_id/contract

    if job_id not exist{
        continue
    }

    job_context = GET /job/$job_id
    input_data = GET /data/$job_context.input_data/blob
    exec ./bin < cat input_data > output_data
    data_id = POST /data output_data
    POST /job/$job_id set status=finished, output_data = data_id
}
```

## worker_id = POST /worker

```
request body
content-type = application/json
{
    "runtime": [
        "default",
        "dlabLua"
    ]
}

response body
content-type = application/json
{
    "code": 0,
    "status": "ok",
    "message": "?",
    "id": "78204124132",
    "runtime": [
        "???",
        "???"
    ]
}
```

## job_id = POST /worker/$worker_id/contract

```
request body
content-type = application/json
{
    "id": "78204124132"
    "tags": [],
    "timeout": 20
}

response body
content-type = application/json
{
    "code": 0,
    "status": "ok",
    "message": "no job",
    "id": "78204124132"
}
```

## job_context = GET /job/$job_id

```
request body
content-type = application/json
{
    "job_id": "78204124132"
}

response body
content-type = application/json
{
    "code": 0,
    "message": "???",
    "id": "78204124132",
    "input": {
        "id": "78204124132"
    },
    "output": {
        "id": "78204124132"
    },
    "lambda": {
        "id": "78204124132",
        "codex": "78204124132",
        "runtime": "default"
    },
    "state": "Running"
}
```

## input_data = GET /data/$data_id/blob

```
response body

if status code == 200
content-type = application/octet-stream
{
    "file goes there"
    "octet-stream means any file (binary or not)"
}

elif status code == 404
content-type = application/json
{
    "code": 601,
    "status": "error",
    "message": "BLOB data object is not exist"
}
```

## data_id = POST /data

```
request body
content-type = multipart/form-data
{
    "file goes there"
    "multipart/form-data means any file (binary of not)"
}

response body
content-type = application/json
{
    "code": 0,
    "status": "ok",
    "message": "http: no such file",
    "id": "78204124132",
    "checksum": "78204124132"
}
```

##
