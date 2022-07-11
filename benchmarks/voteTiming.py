import time
from multiprocessing import Process, Queue

import requests
import json

url = 'http://localhost:8080/user/vote/'
headers = {'Content-type': 'application/json; charset=utf-8'}
invcode = "hpJEsLgp E8aAwnww YZs3BZ0r zLLztv1a kWxezMG1 VWlN2X7d QgYai2yE 7f6mJleA nHt87BCe MsxiGIU8 Bn82aoQ2 ZufPPWv2 ewknl8zF IiSFAQh1 DMlpZmy1 yBL0K5r8 it64MaDn GzUaqdBY UentQzX2 ehFYwMG3 QDOahq80 g29oVy9O 37LTxJV2 RaaqhWrt ryIHMA8s xhIEeMeW sgQeBKk0 OE3vFM0Y OyrvAbqO fxJmPVIV 9x2g3efz CYZ9Z9ew Udp3wecD E4lLtFRK mEcitAH5 7oOl0xaH fboUV9qZ RxWkXrCn Vm6bOHwj c2Rucelk pCx893OZ y1LWQj9Y gyf4Xg13 zkib8Nn3 MwQrVgSG asNeAhgf JXyEPJ2n h4cdYNNS BZVSGbmb l5lyKpvx EMqES5CX DUAPlaYH 8d0xcGmk 8JzX5zVZ Sl6dg5JD xYDp81cN 7B1ppdpz eKYl1XDa d1i1AroQ Cry8SXK3 n590k6i3 QRpR2nf4 xvCEk74h 2eeKAq2U REWLEAsX w1hY8M4w eo08fw1Y iaUfsvX7 I0nLHhnU ZR4KIByD O2f6WiL7 cLQMlYi1 Ax0aBXQ2 gXjoNS3t cJkoiSbH FUAM7sJf ymD2sacV lhjaZntf Zjo7Zqf7 4IvaQITw QhTZz9vO bkK3UXs8 sKfAFqwp 04FH1PbP WJYSr0kw Ct7qEfhn hovydJy3 AMdw14Jr 9tYaJiFG WntOPZsM PFu0EPZ5 Y57ojur0 ui9UxU0f RISFPLQd sJ8yT0on xJGTIWxN kO03zr8V 3QmPjf2s FugLhi8N kDp8Q2yT".split()

def f(i, q):
    body = json.dumps({"inv_code": invcode[i],
                       "event_id": 95,
                       "voted_candidate_names": ["1"]})

    tic = time.perf_counter()
    rep = requests.post(url, headers=headers, data=body)
    toc = time.perf_counter()
    q.put(toc - tic)
    print(f"time: {toc - tic:0.4f}")

    print(rep.json())

if __name__ == "__main__":
    ps = []
    q = Queue()
    tic = time.perf_counter()

    for i in range(0, len(invcode)):
        p = Process(target=f, args=(i,q,))
        p.start()
        ps.append(p)
    for p in ps:
        p.join()
    toc = time.perf_counter()
    print(f"total: {toc - tic:0.4f}")

    while not q.empty():
        print(q.get())
    # body = json.dumps({"inv_code": 'BZDX7MVc',
    #                    "event_id": 85,
    #                    "voted_candidate_names": ["1"]})
    #
    # tic = time.perf_counter()
    # rep = requests.post(url, headers=headers, data=body)
    # toc = time.perf_counter()
    # print(f"{toc - tic:0.4f} seconds")
    #
    # print(rep.json())