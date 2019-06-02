# mapboxcli

The official [Mapbox CLI](https://github.com/mapbox/mapbox-cli-py) is an (almost) complete wrapper for the Mapbox API.

For day-to-day operations, such as changing the underlying dataset for a layer, there is no functionality in the official CLI.

This CLI fills in these gaps and enables deployment procedures which manage Mapbox styles, layers and datasets.

## Getting started

To install, run:

```bash
go get -u github.com/sebnyberg/mapboxcli

mapboxcli --help
```

## Configuration

The Mapbox username and access token can passed either:

1. As commandline flags with --username and --access-token
2. As environment variables named MAPBOX_USERNAME and MAPBOX_ACCESS_TOKEN
3. By adding default flags to `mapbox config set --username $MY_USERNAME --access-token $MY_ACCESS_TOKEN`

Default flags are stored in `~/.mapboxcli/config.yml`. View contents with

```bash
mapbox config show
```

To reset parameters, run

```bash
mapbox config reset
```

## Scenario: renaming the layer in a style

Fetch styles:

```bash
# list styles
mapbox list styles --username myuser --access-token myaccesstoken

# configure
```
