from fastapi import FastAPI, Depends
from sqlalchemy.orm import Session
from db import SessionLocal, engine,SERVER_DB_HOST
import models
import crud

# database
models.Base.metadata.create_all(bind=engine)

app = FastAPI()


# Dependency
async def get_db():
    try:
        db = SessionLocal()
        yield db
    finally:
        db.close()


@app.get('/greet')
async def greet(name: str, db: Session = Depends(get_db)):
    db_query = crud.create_greet(db=db, name=name)
    return {'message': f"Привет, {name} от Python!"}


@app.get('/greet/history')
async def history(db: Session = Depends(get_db)):
    db_query = crud.get_all_greets(db)
    return {'message': db_query}