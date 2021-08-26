import requests
import redis


request = requests.get('http://0.0.0.0')
json = request.json()
print(json)  # {'error': False}

client = redis.Redis(host='localhost', port=6379, password='redis_pass')
client.set('error', str(json.get('error')))
value = client.get('error')
print(value.decode())  # False
