import requests
import os
from dotenv import load_dotenv

register_url = 'https://api2.watttime.org/v2/register'
load_dotenv()

USERNAME = os.getenv('USERNAME')
PASSWORD = os.getenv('PASSWORD')
EMAIL = os.getenv('EMAIL')
ORG = os.getenv('ORG')


params = {'username': USERNAME,
         'password': PASSWORD,
         'email': EMAIL,
         'org': ORG}
rsp = requests.post(register_url, json=params)
print(rsp.text)