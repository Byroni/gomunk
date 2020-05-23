# gomunk

GoMunk is a CLI tool written in Go for uploading and managing files to a personal file store

**Available providers:**

AWS

---

## Installation
Download the latest version of gomunk from the [releases page](https://github.com/Byroni/gomunk/releases).

You must have your AWS credentials setup on your system. You can follow [these instructions](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) to get them ready.

```
Your AWS user must have programmatic access and the proper S3 access policy
```

### Configuration setup
Before you can start using GoMunk, you must set up your configuration file. Create a `config.yml` file in the same directory you have your executable.
If GoMunk cannot find the config file in the executable directory, it will default to the project root. Otherwise, it will exit.

For development, place the config file in the project root folder.

```
~/
  gomunk
  config.yml
```

#### Configuration options
| Name | Default | Options | Description | Required |
| --- | --- | --- | --- | --- |
|  PROVIDER  | | aws |Your configured cloud provider. Gomunk currently only supports AWS | yes | 
| AWS_REGION | us-east-1 | Any valid AWS region | AWS region | no |
| AWS_BUCKET | gomunk-file-store | | AWS bucket to store to be used as your cloud file store| no |

See `config.example.yml` for example usage.

## Usage
Upload a file: 
`gomunk upload /path/to/file.txt`

List files:
`gomunk ls`

Download a file: 
`gomunk get file.txt`
