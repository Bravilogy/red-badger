# Martian Robots
The surface of Mars can be modelled by a rectangular grid around which robots are able to move according to instructions provided from Earth. You are to write a program that
determines each sequence of robot positions and reports the final position of the robot.

### Docker build
To build a docker container
```bash
docker build -t martians .
```

### Run program
To run program, it needs an `input` file that can be specified with `--input` flag. To run the program locally:
```go
go run main.go --input /path/to/input.txt
```

or mount an `input` file in docker container
```
docker run \
	--mount type=bind,source="$(pwd)"/input.txt,target=/tmp/input.txt \
	martians app --input /tmp/input.txt

```

### Run tests
```go
go test ./...
```

Thank you for an awesome challenge.