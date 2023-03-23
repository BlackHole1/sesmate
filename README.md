# SESMATE

## sync

Synchronize local templates to AWS SES service, supporting create, delete, and update operations.

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
      - ses_templates/**.json

  workflow_dispatch:
    inputs:
      tags:
        required: false
        description: ""

jobs:
  migration:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Install sesmate
        run: |
          sudo curl -fsSL -o /usr/local/bin/sesmate https://github.com/BlackHole1/sesmate/releases/latest/download/sesmate-linux-amd64
          sudo chmod +x /usr/local/bin/sesmate

      - name: Run Sync
        run: |
          sesmate sync --dir ./ses_templates --remove
        env:
          AWS_AK: ${{ secrets.AWS_AK }}
          AWS_SK: ${{ secrets.AWS_SK }}
          AWS_REGION: ${{ secrets.AWS_REGION }}
```

#### Use aws credentials

```shell
sesmate sync --dir ./templates
```

#### Use aws profile

```shell
AWS_PROFILE=dev sync --dir ./templates
```

#### Use aws access key and secret key and region

```shell
sesmate sync --dir ./templates --ak AKIAIOSFODNN7EXAMPLE --sk wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY --region us-east-1
```

#### Custom endpoint

```shell
sesmate sync --dir ./templates --endpoint http://localhost:4579
```

#### Use environment variables

```shell
export AWS_AK=AKIAIOSFODNN7EXAMPLE
export AWS_SK=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
export AWS_ENDPOINT=http://localhost:4579
export AWS_REGION=us-east-1
sesmate sync --dir ./templates
```

## zsh completion
```shell
echo "source <(sesmate completion zsh); compdef _sesmate sesmate" >> ~/.zshrc
```
