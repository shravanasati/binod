import requests
import json


def insult(name: str) -> str:
    url = f"https://insult.mattbas.org/api/insult.json?who={name}"
    r = requests.get(url)
    try:
        r.raise_for_status()
        data = json.loads(r.text)
        if data["error"]:
            raise Exception("Error: {}".format(data))

        return data["insult"]

    except Exception as e:
        print(e)
        return f"Trying to roast {name} ruined my internal mechanism."
