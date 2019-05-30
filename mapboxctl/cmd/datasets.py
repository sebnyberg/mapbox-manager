import click


@click.group(short_help="Read and write Mapbox datasets")
@click.pass_context
def datasets(ctx):
    click.echo('Datasets!!')


@datasets.command(short_help="Get a dataset")
@click.pass_context
def get(ctx):
    click.echo('hello, world!')
