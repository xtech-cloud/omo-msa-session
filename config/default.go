package config

const defaultYAML string = `
service:
    address: :7078
    ttl: 15
    interval: 10
logger:
    level: info
    dir: /var/log/msa/
basic:
    timeout: 3600
    secret: "yumei2020-platfom"
`
