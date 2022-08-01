# Govern

<img src="./assets/logo.png" alt="Govenr Logo" width="175"/>

> A CLI tool to validate the commit message based on [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)

## Install

Download the binary from [releases section](https://github.com/shadyaziza/govern/releases), or

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go install github.com/shadyaziza/govern@latest
```

## Example Usage

#### Configure your git hooks to use `govern`

```sh
#!/bin/sh

TEMPORARY_FILE_PATH=$1

COMMIT_MSG=`head -n1 "$TEMPORARY_FILE_PATH"`

if govern --message "$COMMIT_MSG" ;then

    echo  "💪 Valid commit message!  ✔️ 🎯 ✔️ 🎯"
    exit 0
else

	echo "Invalid commit message"
    exit 1
fi

```

## License

```

Copyright 2022 Shady Aziza

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
