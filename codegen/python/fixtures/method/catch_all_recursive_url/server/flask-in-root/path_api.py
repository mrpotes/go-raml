# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.

from flask import Blueprint
import handlers


path_api = Blueprint('path_api', __name__)


@path_api.route('/', defaults={'path': ''}, methods=['GET'], strict_slashes=False)
@path_api.route('/<path:path>', methods=['GET'])
def byPath_get(path):
    """
    It is handler for GET /<path:*>
    """
    return handlers.byPath_getHandler(path)