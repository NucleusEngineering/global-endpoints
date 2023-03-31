build:
	make -C container build

infrastructure:
	make -C infrastructure init apply

.PHONY: build infrastructure