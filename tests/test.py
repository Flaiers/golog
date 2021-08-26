import requests


request = requests.get('http://0.0.0.0')
print(request.json())  # {'error': False}
