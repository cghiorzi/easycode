
from flask import Flask, request, render_template
import os


app = Flask(__name__, static_url_path='/static')

# Determine our running port.  Updated to support Diego which uses the
# 'PORT' environment variable, vs non-Diego which uses VCAP_APP_PORT

if os.getenv('VCAP_APP_PORT'):
    port = os.getenv('VCAP_APP_PORT')
elif os.getenv('PORT'):
    port = os.getenv('PORT')
else:
    port = "8080"


def getName():
	return "Philadelphia, PA"
	
@app.route('/')
def index_page():
    return render_template('index.tmpl',Name=getName())


if __name__ == '__main__':
    app.debug = True
    print "Running on Port: " + port
    app.run(host='0.0.0.0', port=int(port))