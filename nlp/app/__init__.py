from flask import Flask, Blueprint

bp = Blueprint('routes', __name__)

def new_engine():
    engine = Flask(__name__)
    engine.register_blueprint(bp)
    engine.config['JSON_AS_ASCII'] = False
    return engine
from . import routes
from . import snownlp
