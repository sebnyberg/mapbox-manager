import sys
import os

import click

import log
from cmd.datasets import datasets


@click.group()
@click.pass_context
@click.option('--access-token', help="Your Mapbox access token.")
def cli(ctx, access_token):
    """mapboxctl - manage your Mapbox assets"""
    ctx.obj = {}

    access_token = access_token or os.environ.get('MAPBOX_ACCESS_TOKEN')

    log.configure()

    click.echo('Test!')


cli.add_command(datasets)

sys.exit(cli())
