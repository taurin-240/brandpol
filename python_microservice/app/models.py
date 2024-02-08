import datetime

from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy import Column, VARCHAR, DateTime
from sqlalchemy.sql import func
from db import Base
import uuid


class Greet(Base):
    __tablename__ = 'greets'
    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4, unique=True)
    name = Column(VARCHAR(90))
    created_at = Column(DateTime(timezone=True), default=func.now())
    greet = Column(VARCHAR(300), default=f"Привет, {name} от Python!")
