export GO111MODULE=on
export GOSUMDB=off
export GOPROXY=https://mirrors.aliyun.com/goproxy/
go install omo.msa.session
mkdir _build
mkdir _build/bin

cp -rf /root/go/bin/omo.msa.session _build/bin/
cp -rf conf _build/
cd _build
tar -zcf msa.session.tar.gz ./*
mv msa.session.tar.gz ../
cd ../
rm -rf _build
