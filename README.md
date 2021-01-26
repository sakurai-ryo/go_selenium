# Go & Selenium

```bash
.
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── go-stacks                 <-- Source code for a lambda function
│   ├── main.go                 <-- Lambda function code
├── layer                       <-- Selenium & Driver
└── template.yaml
```

### デプロイバケットの作成
```bash
aws s3mb "s3://${STACK_BUCKET}"
```