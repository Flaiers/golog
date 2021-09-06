from dotenv import load_dotenv
from datetime import datetime

import requests
import os


load_dotenv()

data = {
    'date': str(datetime.now()),
    'url': 'https://fla.codes/',
    'method': 'GET',
    'status': 404,
    'user_id': 2,
    'body': None,
    'comment': 'test',
}

request = requests.post('https://logging.fla.codes/', json=data)
json = request.json()
print(json)  # {'error': False, 'data': 'ok'}
