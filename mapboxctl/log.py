import sys
import logging


def configure(log_level=30):
    logging.basicConfig(stream=sys.stderr, level=log_level)
