# local-app

## devtool-cli

All of the accessible commands can be listed with `devtool --help` .

### Installing

1. Download the CLI for your platform and make it executable:

```bash
wget -O devtool https://devtool.io/static/bin/devtool-cli-darwin-arm64
chmod u+x devtool
```

2. Optionally, make it available globally. On macOS:

```bash
sudo mv devtool /usr/local/bin/
```

### Usage

Start by logging in with `devtool login`, which will also create a default context in the configuration file (`~/.devtool/config.yaml`).

### Development

To develop the CLI with Devtool, you can run it just like locally, but in Devtool workspaces, a browser and a keyring are not available. To log in despite these limitations, provide a PAT via the `DEVTOOL_TOKEN` environment variable, or use the `--token` flag with the login command.

#### In a Devtool workspace

[![Open in Devtool](https://www.devtool.io/svg/open-in-devtool.svg)](https://devtool.io/#https://github.com/khulnasoft/devtool)

You will have devtool-cli ready as `devtool` on any Workspace based on `https://github.com/khulnasoft/devtool`.

```
# Reinstall `devtool`
leeway run components/local-app:install-cli

# Reinstall completion
leeway run components/local-app:cli-completion
```

### Versioning and Release Management

The CLI is versioned independently of other Devtool artifacts due to its auto-updating behaviour.
To create a new version that existing clients will consume increment the number in `version.txt`. Make sure to use semantic versioning. The minor version can be greater than 10, e.g. `0.342` is a valid version.

## local-app

**Beware**: this is very much work in progress and will likely break things.

### How to install

```
docker run --rm -it -v /tmp/dest:/out eu.gcr.io/devtool-core-dev/build/local-app:<version>
```

### How to run

```
./local-app
```

### How to run in Devtool against a dev-staging environment

```
cd components/local-app
BROWSER= DEVTOOL_HOST=<URL-of-your-preview-env> go run main.go --mock-keyring run
```
