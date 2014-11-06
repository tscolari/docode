
spec:
	ginkgo -r

get-deps:
	go get gopkg.in/yaml.v2
	go get github.com/onsi/ginkgo
	go get github.com/onsi/gomega
