# omo-msa-session
Micro Service Agent - session

.PHONY: call
call:
	MICRO_REGISTRY=consul micro call omo.msa.session SessionService.Create '{"user":"hzz"}'
	MICRO_REGISTRY=consul micro call omo.msa.session SessionService.CheckAvailable '{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.g93kIDX3blIWfAbREV7Aq87N97rhEyBEhTMH_cxKSn0"}'
