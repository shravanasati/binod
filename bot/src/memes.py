import requests
from bs4 import BeautifulSoup
import random


def get_meme() -> str:
    try:

        subreddit = random.choice(
            [
                "dankmemes",
                "memes",
                "memesoftheyear",
                "programmerhumour",
                "AdviceAnimals",
                "marvelmemes",
            ]
        )
        url = f"https://reddit.com/r/{subreddit}/"
        r = requests.get(url)
        soup = BeautifulSoup(r.content, "html.parser")
        while soup.title.string == "Reddit - Dive into anything":
            subreddit = random.choice(
                ["dankmemes", "memes", "memesoftheyear", "programmerhumour"]
            )
            url = f"https://reddit.com/r/{subreddit}/"
            r = requests.get(url)
            soup = BeautifulSoup(r.content, "html.parser")

        images = []
        for i in soup.find_all("img", {"alt": "Post image"}):
            images.append(i.get("src"))

        print(images)
        return random.choice(images)

    except Exception as e:
        print(e)
        return "Unable to find a meme at the moment! Try again later."


if __name__ == "__main__":
    print(get_meme())
