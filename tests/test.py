from dotenv import load_dotenv
from datetime import datetime

import requests
import psycopg2
import os


load_dotenv()


def send_test(url):
    data = {
        'date': str(datetime.now()),
        'url': 'https://fla.codes/',
        'method': 'GET',
        'status': 404,
        'user_id': 2,
        'body': None,
        'comment': 'test',
    }

    request = requests.post(url, json=data)
    json = request.json()
    print(json)  # {'error': False, 'data': 'ok'}


def psql_test(dsn):
    connection = psycopg2.connect(dsn)
    cursor = connection.cursor()

    query = '''
        SELECT COUNT(*)
        FROM logging;
    '''

    cursor.execute(query)
    for id in cursor.fetchone():
        print(id)

    cursor.close()
    connection.close()


send_test('https://logging.fla.codes/')
psql_test(os.getenv('DB_DSN'))
