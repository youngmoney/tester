#!/usr/bin/env bash

function run() {
	go run . --config tests/simple.config.yaml $@
}

diff <(run test) <(echo pre test; echo complicated command)
diff <(run test -a) <(echo pre test; echo all tests)
diff <(run test a b) <(echo pre test; echo complicated command; echo a; echo b)
diff <(run test -- -a -b) <(echo pre test; echo complicated command; echo -a; echo -b)
diff <(run test -- -a b) <(echo pre test; echo complicated command; echo -a; echo b)
diff <(run test -a -- b) <(echo pre test; echo all tests)
diff <(run test -- -a -- b) <(echo pre test; echo complicated command; echo -a; echo --; echo b)
diff <(run test a -- b) <(echo pre test; echo complicated command; echo a; echo --; echo b)
