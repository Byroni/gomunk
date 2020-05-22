# gomunk

GoMunk is a CLI tool written in Go for uploading and managing files to a personal file store

**Available providers:**

AWS,
Google Cloud Platform

---

## Installation
Download the latest version of gomunk from the [releases page](https://github.com/Byroni/gomunk/releases).

You must have your AWS credentials setup on your system. You can follow [these instructions](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) to get them ready.

## Usage
Upload a file: 
`gomunk upload /path/to/file.txt`

List files:
`gomunk ls`

Download a file: 
`gomunk get file.txt`
