from flask import jsonify, request

from . import users


@users.route('/users', methods=['POST'])
def add_user():
    users = []
    new_user = request.get_json()
    users.append(new_user)
    return jsonify({
        "users": users,
        "message": "User added successfully",
        "status": "success",
        "status_code": 201
    })