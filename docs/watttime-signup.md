# Signing up to Watttime


## Python Tooling

The tooling uses Python, and we recommend using a conda environment when installing requirements, for example:

```bash
$ cd wattime
$ conda create --name watttime python=3.9 -y
$ conda activate watttime
$ pip install -r requirements-dev.txt
```

Edit the `.env.sample` file and save it as `.env` with your updated parameters, and then run `watttime.py`.

### Getting your Watttime key

Source: https://www.watttime.org/api-documentation/#authentication


```python
import requests
register_url = 'https://api2.watttime.org/v2/register'
params = {'username': ‘freddo’,
         'password': 'the_frog',
         'email': 'freddo@frog.org',
         'org': 'freds world'}
rsp = requests.post(register_url, json=params)
print(rsp.text)
```
