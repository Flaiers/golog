from dotenv import load_dotenv
from datetime import datetime

import requests
import redis
import os


load_dotenv()

data = {
    'date': str(datetime.now()),  # 2021-08-28 00:50:11.331569
    'url': 'https://google.com/',
    'method': 'GET',
    'status': 200,
    'user': 'bigin.maks@gmail.com',
    'headers': '',
    'body': '',
    'comment': '',
}

request = requests.post('https://analytics.fla.codes/', json=data)
json = request.json()
print(json)  # {'error': False, 'data': 'ok'}

client = redis.Redis(host='localhost', port=6379, db=0, password=os.environ.get('REDIS_PASSWORD'))
value = client.get(data.get('date',))
print(value.decode())  # "date":"2021-08-28 00:50:11.331569","url":"https://google.com/","method":"GET","status":200,"user":"bigin.maks@gmail.com","headers":"","body":"","comment":""}
