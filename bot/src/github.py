import requests
import json


def github_user(username: str) -> str:
    url = f"https://api.github.com/users/{username}"
    r = requests.get(url)
    try:
        r.raise_for_status()
        data = json.loads(r.text)
        answer_text = ""
        answer_text += "\n"
        answer_text += "Name: {}\n".format(data["name"])
        answer_text += "Avatar: {}\n".format(data["avatar_url"])
        answer_text += "Followers: {}\n".format(data["followers"])
        answer_text += "Following: {}\n".format(data["following"])
        answer_text += "Repos: {}\n".format(data["public_repos"])
        answer_text += "Bio: {}\n".format(data["bio"])
        return answer_text

    except Exception as e:
        print(f"There was a problem: {e}")
        return "Unable to find a GitHub user named " + username


def github_repo(username: str, repo: str) -> str:
    url = f"https://api.github.com/repos/{username}/{repo}"
    r = requests.get(url)
    try:
        r.raise_for_status()
        data = json.loads(r.text)
        answer_text = ""
        answer_text += "Repo name: {}\n".format(data["name"])
        answer_text += "Owner name: {}\n".format(data["owner"]["login"])
        answer_text += "Language: {}\n".format(data["language"])
        answer_text += "Description: {}\n".format(data["description"])
        answer_text += "Stars: {}\n".format(data["stargazers_count"])
        answer_text += "Watchers: {}\n".format(data["watchers_count"])
        answer_text += "Forks: {}\n".format(data["forks_count"])
        answer_text += "License: {}\n".format(data["license"]["name"])
        answer_text += "URL: {}\n".format(data["html_url"])
        return answer_text

    except Exception as e:
        print(f"There was a problem: {e}")
        return f"Unable to find a GitHub repo named {username}/{repo}."


if __name__ == "__main__":
    print(github_user("Shravan-1908"))
    print(github_repo("Shravan-1908", "hydra"))
