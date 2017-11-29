# Domain Expert Bot - Stackoverflow Poller
Crawler for Kubernetes & Istio tagged questions on stackoverflow

## Installation

- glide install -v
- go build .
- docker build -t <image_name> .
- docker push <image_name>
- kubectl create -f cronjob.yaml