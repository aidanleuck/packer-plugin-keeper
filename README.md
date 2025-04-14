# Keeper Packer Plugin

This repository produces a Packer plugin which allows users to load secrets from Keeper Secrets Manager. 

## Quick Start

See the following [user facing documentation](https://github.com/aidanleuck/packer-plugin-keeper/blob/main/docs/README.md) 
## Build from source

1. Clone this GitHub repository locally.

2. Run this command from the root directory: 
```shell 
go build -ldflags="-X github.com/aidanleuck/packer-plugin-keeper/version.VersionPrerelease=dev" -o packer-plugin-scaffolding
```

3. After you successfully compile, the `packer-plugin-scaffolding` plugin binary file is in the root directory. 

4. To install the compiled plugin, run the following command 
```shell
packer plugins install --path packer-plugin-keeper github.com/aidanleuck/keeper
```

### Build on *nix systems
Unix like systems with the make, sed, and grep commands installed can use the `make dev` to execute the build from source steps. 

### Build on Windows Powershell
The preferred solution for building on Windows are steps 2-4 listed above.
If you would prefer to script the building process you can use the following as a guide

```powershell
$MODULE_NAME = (Get-Content go.mod | Where-Object { $_ -match "^module"  }) -replace 'module ',''
$FQN = $MODULE_NAME -replace 'packer-plugin-',''
go build -ldflags="-X $MODULE_NAME/version.VersionPrerelease=dev" -o packer-plugin-scaffolding.exe
packer plugins install --path packer-plugin-scaffolding.exe $FQN
```

## Running Acceptance Tests

Make sure to install the plugin locally using the steps in [Build from source](#build-from-source).

Once everything needed is set up, run:
```
PACKER_ACC=1 go test -count 1 -v ./... -timeout=120m
```

This will run the acceptance tests for all plugins in this set.

# Requirements

-	[packer-plugin-sdk](https://github.com/hashicorp/packer-plugin-sdk) >= v0.5.2
-	[Go](https://golang.org/doc/install) >= 1.20
