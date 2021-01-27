package config

const defaultYAML string = `
service:
    address: :7089
    ttl: 15
    interval: 10
logger:
    level: info
    dir: /var/log/msa/
basic:
    timeout: 43200
    secret: "yumei2020-platfom"
`
