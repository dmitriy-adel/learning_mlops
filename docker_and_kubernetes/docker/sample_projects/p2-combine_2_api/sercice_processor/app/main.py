from fastapi import FastAPI
from pydantic import BaseModel

app_processor = FastAPI(title="Text Processing API", version="1.0.0")


class ProcessReques(BaseModel):
    text: str


@app_processor.post("/get_total_alpha")
def get_total_alpha(request: ProcessReques):
    return sum(t.isalpha() for t in request.text)
    
@app_processor.post("/get_total_digits")
def get_total_digits(request: ProcessReques):
    return sum(t.isdigit() for t in request.text)
