import os, logging
from . import create_app
from flask import jsonify, request, abort


app = create_app()

# Auth middleware
@app.before_request
def check_auth():
    if request.endpoint != 'home' and not request.headers.get('Authorization'): # or (and not request.headers.get('X-API-Key'))
        abort(401, description="Unauthorized")

logging.basicConfig(level=logging.INFO)
@app.after_request
def log_response(response):
    logging.info(f"{request.method} {request.path} - {response.status_code}")
    return response

@app.route("/")
def home():
    return jsonify({"message": "Hello, World!"})

if __name__ == '__main__':
    if os.environ.get("FLASK_ENV") == "production":
        app.run(host="0.0.0.0", port=80)
    else:
        app.run(debug=True) 