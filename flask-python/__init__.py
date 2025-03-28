from flask import Flask

def create_app():

    app = (Flask(__name__))

    from .routes.users import users
    app.register_blueprint(users)

    return app