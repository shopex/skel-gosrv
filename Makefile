OBJ := @{SKEL}

all: ${OBJ}

${OBJ}: docs.go *.go */*.go
	go build .

docs.go: controller/*.go
	lsv generate docs

start: all
	./${OBJ} server start

clean:
	-rm ${OBJ} docs.go

test:
	(cd testcase ; go test)
