from flask import jsonify
from flask.views import MethodView
from sqlalchemy import text
# from .add_user import execute_and_serialize  (had an error of circular imports)
from ... import db

def execute_and_serialize(query, params=None):
    """Execute raw SQL and return serialized results"""
    result = db.session.execute(text(query), params or {})
    column_names = result.keys()
    
    return [dict(zip(column_names, row)) for row in result]

class UserController(MethodView):
    # db_users = []
    def __init__(self):
        pass

    def get(self):
        self.db_users = execute_and_serialize("SELECT * FROM user")
        return jsonify({
            "users": self.db_users,
            "message": "Users retrieved successfully",
            "status": "success",
            "status_code": 200
        }), 200

# Define the endpoint
from . import users
users.add_url_rule('/users', view_func=UserController.as_view('users'))