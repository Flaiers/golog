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
print(json)  # {'error': False, 'data': 'ok'}

client = redis.Redis(host='localhost', port=6379, password=os.environ.get('REDIS_PASSWORD'))
value = client.get('data')
print(value.decode())  # {"date":"","url":"https://google.com/","method":"GET","status":200,"user":"bigin.maks@gmail.com","headers":"","body":"","comment":""}