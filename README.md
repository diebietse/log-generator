# Log Generator

A very basic program that reads a list of lines to log from a config file.

It then logs the lines in the file one at a time and then loops back and repeats it

## Run Locally

You can run it locally with:

```console
go run main.go
```

Or with a custom delay between logs and a custom config with:

```console
go run main.go --log_source_file_path=example_logs.txt --log_delay_seconds=1
```

You can also run it in docker using:

```console
make docker
make run docker
```
