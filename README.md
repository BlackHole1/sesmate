# SESMATE

## Install

```shell
curl -fsSL https://raw.githubusercontent.com/BlackHole1/sesmate/main/install.sh | sh -
# or
wget -qO- https://raw.githubusercontent.com/BlackHole1/sesmate/main/install.sh | sh -
```

## sync

Synchronize local templates to AWS SES service, supporting create, delete, and update operations.

> AWS SES Template Structure: https://docs.aws.amazon.com/ses/latest/dg/send-personalized-email-api.html#send-personalized-email-create-template

### Command Line Options

- `--dir` - The local template directory.
- `--remove` - Delete remote template when it is not found locally.
- `--ak` - The AWS access key.
- `--sk` - The AWS secret key.
- `--endpoint` - The AWS endpoint.
- `--region` - The AWS region.
- `--help` - Print usage help.

### Usage

#### GitHub Action

```yaml
name: "Sync SES Template"
on:
  push:
    branches:
      - main
    paths:
      - sestemplate/**.json

  workflow_dispatch:
    inputs:
      tags:
        required: false
        description: ""

jobs:
  sync:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Install sesmate
        run: |
          sudo curl -fsSL -o /usr/local/bin/sesmate https://github.com/BlackHole1/sesmate/releases/latest/download/sesmate-linux-amd64
          sudo chmod +x /usr/local/bin/sesmate

      - name: Run Sync
        run: sesmate sync --dir ./sestemplate --remove
        env:
          AWS_AK: ${{ secrets.AWS_AK }}
          AWS_SK: ${{ secrets.AWS_SK }}
          AWS_REGION: ${{ secrets.AWS_REGION }}
```

or

```yaml
# ...
- name: Configure AWS Credentials
  uses: aws-actions/configure-aws-credentials@v1
  with:
    aws-region: ${{ secrets.AWS_REGION }}
    aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
    aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
    role-to-assume: ${{ secrets.AWS_ROLE_TO_ASSUME }}
    role-duration-seconds: 120
    role-session-name: ${{ secrets.AWS_ROLE_SESSION_NAME }}
# ...
- name: Run Sync
  run: sesmate sync --dir ./sestemplate --remove
```

#### Use aws credentials

```shell
sesmate sync --dir ./sestemplate
```

#### Use aws profile

```shell
AWS_PROFILE=dev sync --dir ./sestemplate
```

#### Use aws access key and secret key and region

```shell
sesmate sync --dir ./sestemplate --ak AKIAIOSFODNN7EXAMPLE --sk wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY --region us-east-1
```

#### Custom endpoint

```shell
sesmate sync --dir ./sestemplate --endpoint http://localhost:4579
```

#### Use environment variables

```shell
export AWS_AK=AKIAIOSFODNN7EXAMPLE
export AWS_SK=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
export AWS_ENDPOINT=http://localhost:4579
export AWS_REGION=us-east-1
sesmate sync --dir ./sestemplate
```

## gen

Generate go file from template files.

### Command Line Options

- `--dir` - Local template directory.
- `--output` - Output directory. Default: "./sestemplate".
- `--filename` - Output file name. Default: "name".
- `--package-name` - GO package name. Default: "sestemplate".
- `--prefix` - Prefix of generated const. Default: "".
- `--case` - Case of generated const, allowed values: *lower*, *upper*, *camel*, *pascal*, *snake*, *screaming_snake*, *capitalized_snake*. Default: "pascal".
- `--help` - Print usage help.

### Usage

```shell
sesmate gen --dir ./sestemplate --output ./sestemplate --filename name --package-name sestemplate --prefix Name --case pascal
```

## Example

./sestemplate/a.json:
```json
{
    "Template": {
        "TemplateName": "v1_user_info",
        "SubjectPart": "Your Info",
        "TextPart": "Name: {{name}}, age: {{ Age }}, email: {{_email}}, nickname: {{_name}}"
    }
}
```

execute:
```shell
sesmate gen --dir ./sestemplate --output ./sestemplate --filename name --package-name sestemplate --case pascal
```

output file: ./sestemplate/name.go:
```go
package sestemplate

const V1UserInfo = "v1_user_info"

type V1UserInfoData struct {
	Name  string `json:"name"`
	Age   int    `json:"Age"`
	Email string `json:"_email"`
	Name2 string `json:"_name"`
}
```

## zsh completion
```shell
echo "source <(sesmate completion zsh); compdef _sesmate sesmate" >> ~/.zshrc
```
