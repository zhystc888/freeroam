# 根目录 Makefile（freeroam/Makefile）

# 预留服务列表：后续加 ort system 等，按 app/<name> 目录名填写
# SERVICES := gateway
SERVICES := gateway org system

ENV ?=
TAG ?=

.PHONY: help services build image.push $(addsuffix .build,$(SERVICES)) $(addsuffix .image.push,$(SERVICES))

help:
	@echo "Usage:"
	@echo "  make build                 # build all services"
	@echo "  make image.push            # build+push images for all services"
	@echo "  make image.push ENV=uat    # push uat-latest + uat-<tag>"
	@echo "  make image.push TAG=v1.0.0 # override tag"
	@echo ""
	@echo "Single service:"
	@echo "  make gateway.build"
	@echo "  make gateway.image.push ENV=uat"

services:
	@echo "$(SERVICES)"

# build all
build:
	@set -e; \
	for s in $(SERVICES); do \
		echo "==> [$$s] make build"; \
		$(MAKE) -C app/$$s build; \
	done

# image push all (calls each service's image.push)
image.push:
	@set -e; \
	for s in $(SERVICES); do \
		echo "==> [$$s] make image.push ENV=$(ENV) TAG=$(TAG)"; \
		$(MAKE) -C app/$$s image.push ENV="$(ENV)" TAG="$(TAG)"; \
	done

# -------- per-service shortcuts --------
# e.g. make gateway.build
$(addsuffix .build,$(SERVICES)):
	$(MAKE) -C app/$(basename $@) build

# e.g. make gateway.image.push ENV=uat TAG=...
$(addsuffix .image.push,$(SERVICES)):
	$(MAKE) -C app/$(basename $@) image.push ENV="$(ENV)" TAG="$(TAG)"
