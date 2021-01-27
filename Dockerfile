FROM alpine:3.11
ADD omo.msa.session /usr/bin/omo.msa.session
ENV MSA_REGISTRY_PLUGIN
ENV MSA_REGISTRY_ADDRESS
ENTRYPOINT [ "omo.msa.session" ]
