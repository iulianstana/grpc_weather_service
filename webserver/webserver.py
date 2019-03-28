import os
import sys

tmp_path = os.path.join(os.path.dirname(os.path.realpath(__file__)), '..')
if sys.path[0] != tmp_path:
    while tmp_path in sys.path:
        sys.path.remove(tmp_path)
    sys.path.insert(0, tmp_path)

from flask import Flask, jsonify
from flask import request
from flask import abort
from flask import make_response

from clients.weather_service.python.weather_client import WeatherClient


app = Flask(__name__)


@app.errorhandler(404)
def not_found(error):
    """ Handler 404 error in a better way """
    return make_response(jsonify({'error': 'Not found'}), 404)


@app.errorhandler(400)
def not_found(error):
    """ Handler 400 error in a better way """
    return make_response(jsonify({'error': 'Bad request'}), 400)


@app.route('/weather/', methods=['GET'])
def get_weather():
    """
    Use GET method to get today weather value for a given city
    """
    city = request.args.get('city')

    if city is None:
        abort(400, "Need to provide a city.")

    # get weather
    client = WeatherClient()

    current_conditions = client.get({'city': city})

    return jsonify(current_conditions)


if __name__ == '__main__':
    print "Start flask server"
    app.run(host="0.0.0.0", port=5000, debug=True)
