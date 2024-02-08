from sqlalchemy.orm import Session
from sqlalchemy.sql import func
from sqlalchemy import desc
import models
import uuid


def create_greet(db: Session, name: str):
    db_greet = models.Greet(id=uuid.uuid4(), name=name, created_at=func.now(),greet=f"Привет, {name} от Python!")
    db.add(db_greet)
    db.commit()
    db.refresh(db_greet)
    return db_greet


def get_all_greets(db: Session):
    return db.query(models.Greet).order_by(desc(models.Greet.created_at)).all()
