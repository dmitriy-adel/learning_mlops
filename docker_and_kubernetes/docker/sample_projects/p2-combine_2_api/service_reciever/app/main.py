from fastapi import FastAPI
from pydantic import BaseModel
import requests

import os
env_vars = os.environ

app_receiver = FastAPI(title="Text Processing API", version="1.0.0")


class UserReques(BaseModel):
    text: str
    process_type: int


@app_receiver.post("/process_request")
def process_request(request: UserReques):
    payload: dict[str, str] = {"text": request.text}
    endpoint: str = "get_total_alpha" if request.process_type == 0 else "get_total_digits"

    response = send_request(endpoint=endpoint, payload=payload)

    return response


def send_request(endpoint: str, payload: dict):
    url = f"{env_vars['PROCESSOR_HOST']}/{endpoint}"
    try:
        resp = requests.post(url, json=payload)
        resp.raise_for_status()
        return resp.json()
    
    except requests.RequestException as e:
        return {"error": str(e)}
