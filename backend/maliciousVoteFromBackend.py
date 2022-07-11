import requests
import json

url = 'http://localhost:8080/user/vote/'
headers = {'Content-type': 'application/json; charset=utf-8'}

if __name__ == "__main__":
    body = json.dumps({"inv_code": "5HJwsbo3",
                    "event_id": 99,
                    "voted_candidate_names": ["Alice"]})

    rep = requests.post(url, headers=headers, data=body)

    print(rep)