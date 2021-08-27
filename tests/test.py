from dotenv import load_dotenv

import requests
import redis
import os


load_dotenv()

data = {
    'url': 'https://google.com/',
    'method': 'GET',
    'status': 200,
    'user': 'bigin.maks@gmail.com'
}

request = requests.post('https://analytics.fla.codes/', json=data)
json = request.json()
print(json)  # {'error': False, 'data': 'OK'}

client = redis.Redis(host='localhost', port=6379, password=os.environ.get('REDIS_PASSWORD'))
client.set('data', json.get('data'))
value = client.get('data')
print(value.decode())  # OK
