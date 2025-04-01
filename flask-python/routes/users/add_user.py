from flask import jsonify, request
from sqlalchemy import text

from . import users
from . import get_users
from ... import db

# check if user exists
def user_exists(id, name=""):
    """Check if user exists in the database"""
    query = "SELECT COUNT(*) FROM user WHERE id = :id OR name = :name"
    result = db.session.execute(text(query), {"id": id, "name": name}).scalar()
    return result > 0 # True if user exists

def execute_and_serialize(query, params=None):
    """Execute raw SQL and return serialized results"""
    result = db.session.execute(text(query), params or {})
    column_names = result.keys()
    
    return [dict(zip(column_names, row)) for row in result]



@users.route('/users', methods=['POST'])
def add_user():
    new_user = request.get_json()
    
    if user_exists(new_user["id"], new_user["name"]):
        return jsonify({
            "message": "User already exists",
            "status": "error",
            "status_code": 409
        }), 409
    
    try:
        db.session.execute(text("INSERT INTO user (id, name) VALUES (:id, :name)"), {"id": new_user["id"], "name": new_user["name"]})
        db.session.commit()
        # results = db.session.execute(text("SELECT * FROM user")).all()
        all_users = execute_and_serialize("SELECT * FROM user")
        return jsonify({
            "users": all_users,
            "message": "User added successfully",
            "status": "success",
            "status_code": 201
        }), 201
    except Exception as e:
        db.session.rollback()
        return jsonify({
            "message": f"Failed to add user: {str(e)}",
            "status": "error",
            "status_code": 500
        }), 500