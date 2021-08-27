from dotenv import load_dotenv

import requests
import redis
import os


load_dotenv()

request = requests.get('http://analytics.fla.codes/')
json = request.json()
print(json)

client = redis.Redis(host='localhost', port=6379, password=os.environ.get('REDIS_PASSWORD'))
client.set('data', json.get('data'))
value = client.get('data')
print(value.decode())
