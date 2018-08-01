OBJS = fisher.go utils.go
build:
	go build ${OBJS} && mv ./fisher bin