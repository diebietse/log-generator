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

## Run K8s

Create a configmap named `log-config` with key `log-config` containing the example logs:

```console
kubectl create configmap log-config --from-file=log-config=example_logs.txt --dry-run=client -o yaml | kubectl apply -f -
```

Create a deployment of the log-generator:

```console
kubectl apply -f ./k8s.yml
```
