import requests
import json

url = 'http://localhost:8080/user/vote-details/'
headers = {'Content-type': 'application/json; charset=utf-8'}

if __name__ == "__main__":
    body = json.dumps({"inv_code": "xyin68@gatech.edu"})

    rep = requests.post(url, headers=headers, data=body)

    print(rep)