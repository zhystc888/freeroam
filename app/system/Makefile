ROOT_DIR    := $(CURDIR)
NAMESPACE   ?= default
DEPLOY_NAME ?= template-mono-app
DOCKER_NAME ?= system

ENV ?=        # 不传即生产：tag 不带 env 前缀
TAG ?=        # 可选：手动覆盖构建标签，如 TAG=v1.2.3

include ./hack/hack-cli.mk
include ./hack/hack.mk

.PHONY: image.push

image.push: build
	$(eval _TAG := $(shell git rev-parse --short HEAD 2>/dev/null || echo "nogit"))
ifneq (, $(shell git status --porcelain 2>/dev/null))
	$(eval _TAG := $(_TAG).dirty)
endif
	$(eval _TAG := $(if $(strip $(TAG)),$(TAG),$(_TAG)))
	$(eval _TAG_PREFIX := $(if $(strip $(ENV)),$(ENV)-,))

	@echo "==> gf docker push: $(DOCKER_NAME):$(_TAG_PREFIX)latest"
	@gf docker -p -tn $(DOCKER_NAME):$(_TAG_PREFIX)latest

	@echo "==> gf docker push: $(DOCKER_NAME):$(_TAG_PREFIX)$(_TAG)"
	@gf docker -p -tn $(DOCKER_NAME):$(_TAG_PREFIX)$(_TAG)
