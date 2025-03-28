from flask import jsonify

from . import users

@users.route('/users', methods=['GET'])
def get_users():
    users = []
    return jsonify({
        "users": users,
        "message": "Users retrieved successfully",
        "status": "success",
        "status_code": 200
    })