# Go & Selenium

```bash
.
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── go-stacks                 <-- Source code for a lambda function
│   ├── main.go                 <-- Lambda function code
├── layer                       <-- Selenium & Driver
│   ├── unzip                 <-- binary
│   ├── zip                 <-- zipped binary
└── template.yaml
```

## デプロイバケットの作成
```bash
$ aws s3 mb "s3://layer/zip/layer.zip"
$ aws s3 mb "s3://${STACK_BUCKET}"
```
あとで自動化とsamに組み込む

## Layerアップロード
```bash
$ aws s3 sync layer/zip/ s3://selenium-stack/
```
この辺りもあとでsamに組み込み

## 各ブラウザ関連のダウンロード
### chromedriver
```bash
$ curl -SL https://chromedriver.storage.googleapis.com/2.37/chromedriver_linux64.zip > layer/unzip/chromedriver.zip
$ unzip layer/unzip/chromedriver.zip
$ rm layer/unzip/chromedriver.zip
```
### chromium
```bash
serverless-chromium
$ curl -SL https://github.com/adieuadieu/serverless-chrome/releases/download/v1.0.0-37/stable-headless-chromium-amazonlinux-2017-03.zip > layer/unzip/headless-chromium.zip
$ unzip layer/unzip/headless-chromium.zip
$ rm layer/unzip/headless-chromium.zip
```