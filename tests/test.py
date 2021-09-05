from dotenv import load_dotenv
from datetime import datetime

import requests
import os


load_dotenv()

data = {
    'date': str(datetime.now()),  # 2021-08-28 00:50:11.331569
    'url': 'https://google.com/',
    'method': 'GET',
    'status': 404,
    'user_id': 2,
    'headers': '',
    'body': '',
    'comment': '',
}

request = requests.post('https://logging.fla.codes/', json=data)
json = request.json()
print(json)  # {'error': False, 'data': 'ok'}
