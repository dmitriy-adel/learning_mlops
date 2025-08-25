from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI(title="Text Processing API", version="1.0.0")


class TextRequest(BaseModel):
    text: str


@app.post("/count_words")
def count_words(request: TextRequest):
    words = request.text.split()
    return {"word_count": len(words)}


@app.post("/count_chars")
def count_chars(request: TextRequest):
    chars = len(request.text.replace(" ", ""))
    return {"char_count": chars}


@app.post("/reverse")
def reverse_text(request: TextRequest):
    return {"reversed": request.text[::-1]}
