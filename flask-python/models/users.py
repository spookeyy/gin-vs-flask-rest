from sqlalchemy import Column, Integer, String
from .. import db

class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(50), nullable=False)

    def to_dict(self):
        """convert model instance to a dictionary"""
        return {
            "id": self.id,
            "name": self.name
        }
    
    @staticmethod
    def from_dict(data):
        """create a new user instance from dictionary data"""
        return User(
            id=data["id"],
            name=data["name"]
        )