BASED_URL := localhost:9400/api/
JSON_CONTENT_TYPE := --header 'Content-Type: application/json'

.PHONY: post
post: endpoint data clear
	@curl -X POST --silent --location $(BASED_URL)$(e) $(JSON_CONTENT_TYPE) --data '$(shell echo '$(d)' | jq -c .)' | jq


.PHONY: get
get: endpoint clear
	@curl -X GET --silent --location $(BASED_URL)$(e) | jq



endpoint:
ifndef e
	@$(error param 'e' is required [endpoint])
endif

data:
ifndef d
	@$(error param 'd' is required [data])
endif



XDOTOOL := xdotool
CLEAR_TERM_COMB := alt+BackSpace

.PHONY: clear
clear:
	@${XDOTOOL} key ${CLEAR_TERM_COMB}



.PHONY: run
run: clear
	@go run cmd/main.go -env testing