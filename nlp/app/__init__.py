from flask import Flask, Blueprint

bp = Blueprint('app', __name__)


def new_engine():
    engine = Flask(__name__)
    engine.json.ensure_ascii = False
    engine.register_blueprint(bp)
    return engine


from . import routes
from . import snownlp
