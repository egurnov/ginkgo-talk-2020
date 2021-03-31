.PHONY: present
present:
	present -notes

.PHONY: test
test:
	go test -count=1 ./...

.PHONY: ginkgo
ginkgo:
	ginkgo -skip Flaky -skipMeasurements -r

.PHONY: gomega
gomega:
	ginkgo -focus Gomega ./ginkgo

.PHONY: gomega_standalone
gomega_standalone:
	go test -count=1 ./gomega

.PHONY: flaky
flaky:
	ginkgo -untilItFails -focus Flaky ./ginkgo

.PHONY: watch
watch:
	ginkgo watch -notify -r -skip Flaky -skipMeasurements

.PHONY: measure
measure:
	ginkgo -focus Measurements ./ginkgo