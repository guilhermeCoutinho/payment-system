TESTABLE_PACKAGES=`go list ./... | grep 'controller'`

deps:
	@sh ./dev/deps.sh

run:
	@go run *.go serve


.PHONY: mocks
mocks:
	@python3 scripts/generate_mocks.py

unit:
	@go test -v ${TESTABLE_PACKAGES} -tags=unit -coverprofile=unit.coverprofile -count=1