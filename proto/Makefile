PROTOS := $(wildcard */*.proto) 
PROTO_OBJS := $(PROTOS:.proto=.pb.go) 

all: $(PROTO_OBJS)

%.pb.go: %.proto $(PROTOS)
	test -d $(@D) || mkdir $(@D)
	protoc --go_out=../../../.. $<

