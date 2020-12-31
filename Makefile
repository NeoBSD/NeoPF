BINARY_DIR=bin
BINARY_NAME=neopf

DEV_IP=10.23.0.103

.PHONY: all
all: native

${BINARY_DIR}:
	mkdir -p ${BINARY_DIR}

.PHONY: native
native: ${BINARY_DIR}
	go build -o ${BINARY_DIR}/${BINARY_NAME}

.PHONY: freebsd
freebsd: ${BINARY_DIR}
	GOOS=freebsd GOARCH=arm64 go build -o ${BINARY_DIR}/${BINARY_NAME}-freebsd_arm

.PHONY: copy
copy:
	scp ${BINARY_DIR}/${BINARY_NAME}-freebsd_arm ${DEV_IP}:~/${BINARY_NAME}

.PHONY: run
run:
	@ssh ${DEV_IP} ~/${BINARY_NAME}

.PHONY: clean
clean:
	rm -rf ${BINARY_DIR}

