from flask import Blueprint

users = Blueprint('users', __name__)

from . import get_users, add_user