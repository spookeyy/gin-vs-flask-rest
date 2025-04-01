import os

class Config():
    DEBUG = True
    SECRET_KEY = os.environ.get("SECRET_KEY") or "my_precious_secret_key"
    BASE_DIR = os.path.abspath(os.path.dirname(__file__))  # for sqlite
    DB_PATH = os.path.join(BASE_DIR, 'instance' "users.db") # for sqlite

    os.makedirs(os.path.dirname(DB_PATH), exist_ok=True) # 

    SQLALCHEMY_DATABASE_URI = f"sqlite:///{DB_PATH}"

    # postgresql
    # SQLALCHEMY_DATABASE_URI = "postgresql://postgres:spookie@localhost:5432/postgres"
    # postgres is the database name, postgres is the username, localhost is the host, 5432 is the port, users is the schema
    SQLALCHEMY_TRACK_MODIFICATIONS = False