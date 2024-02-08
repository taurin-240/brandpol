#~/python_microservice/app/api/db.py
import socket

from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from setting import USER_POSTGRES,PASS_POSTGRES,DATABASE_NAME,SERVER_DB_HOST


DATABASE_URI = f"postgresql://{USER_POSTGRES}:{PASS_POSTGRES}@{SERVER_DB_HOST}:5434/{DATABASE_NAME}"

engine = create_engine(DATABASE_URI, connect_args={},future=True)

SessionLocal = sessionmaker(autocommit=False, bind=engine,future=True)

Base = declarative_base()

