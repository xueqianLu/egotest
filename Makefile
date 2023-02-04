DEST=${PWD}
BIN=egotest

.PHONY: all bin run sim clean

all: bin

bin:
	ego-go build -o=${BIN}
	ego sign ${BIN}

sim:
	OE_SIMULATION=1 ego run ${BIN}

run:
	ego run ${BIN}

clean:
	@rm -f $(BIN)
