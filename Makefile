
BUILD		:= release
VERSION		:= v1.0.0
REVISION	:= $(shell git rev-parse --short HEAD)

SRCS		:= $(shell find ./src -type f -name '*.go')
VSRCS		:= $(shell find ./vendor -type f -name '*.go')
VSRC_EXIST	:= $(shell find ./vendor -name src)
# LDFLAGS		:= -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
LDFLAGS		:= -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

define vendor_restore
	@gb vendor restore 
endef

BuildDeb = mkdir -p work/opt/bto_ir_advanced_cmd_go; \
		   mkdir -p work/DEBIAN ; \
		   cp ./bin/$(1) ./work/opt/bto_ir_advanced_cmd_go/bto_ir_cmd ; \
		   cp ./config.json ./work/opt/bto_ir_advanced_cmd_go; \
		   cp ./deb/post* ./work/DEBIAN/ ; \
		   sed -e 's/%%DBG%%/$(2)/' -e 's/%%ARCH%%/$(3)/' ./deb/conf > ./work/DEBIAN/control ; \
		   fakeroot dpkg-deb --build ./work . ; \
		   rm -rf ./work

BuildCommand = GOARCH=$(1) GOOS=linux gb build -tags '$(2)' $(3) all 

ifeq ($(BUILD),debug)
TAGS	:= debug
IS_DBG	:= -debug
else
TAGS	:= 
IS_DBG	:= 
endif

.PHONY: all
all: amd64

.PHONY: clean
clean:
	@rm -rf bin
	@rm -rf pkg
	@rm -rf *.deb

.PHONY: allclean
allclean:
	@rm -rf bin
	@rm -rf pkg
	@rm -rf vendor/src
	@rm -rf *.deb

386: $(SRCS)
	$(if $(VSRC_EXIST) ,,$(vendor_restore))
	$(call BuildCommand,386,$(TAGS),$(LDFLAGS))

arm: $(SRCS)
	$(if $(VSRC_EXIST) ,,$(vendor_restore))
	$(call BuildCommand,arm,$(TAGS),$(LDFLAGS))

amd64: $(SRCS)
	$(if $(VSRC_EXIST) ,,$(vendor_restore))
	$(call BuildCommand,amd64,$(TAGS),$(LDFLAGS))

deb-386: ./bin/main-linux-386$(IS_DBG)
	$(call BuildDeb,main-linux-386$(IS_DBG),$(IS_DBG),i386)

deb-amd64: ./bin/main-linux-amd64$(IS_DBG)
	$(call BuildDeb,main-linux-amd64$(IS_DBG),$(IS_DBG),amd64)

deb-armhf: ./bin/main-linux-arm$(IS_DBG)
	$(call BuildDeb,main-linux-arm$(IS_DBG),$(IS_DBG),armhf)

