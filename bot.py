import json
import os
import random
import subprocess
import discord
import requests
import wikipedia
from bs4 import BeautifulSoup
from src.insult import insult
from src.github import github_user, github_repo


client = discord.Client()


@client.event
async def on_ready():
    print('Bot has logged in as {0.user}'.format(client))


@client.event
async def on_message(message):
    command = message.content.lower()

    if message.author == client.user:
        return

    help = "Hello, I am Binod, a bot made by Shravan. The following are the commands:\n\n binod.help = Show this help message\n\n binod.binod = BINOD\n\n binod.news = Top news headlines of the moment\n\n binod.weather <place> = Real time weather of <place> \n\n binod.wiki <search> = Searches the wikipedia for <search> \n\n binod.joke = Get a random joke \n\n binod.corona = Get live status of confirmed corona cases in India \n\n binod.meme = Get a random meme \n\n binod.suggest <suggestion> = Suggest a new feature to be added to the bot \n\n"

    if command.startswith('binod.help'):
        await message.channel.send(help)

    elif command.startswith('binod.binod'):
        await message.channel.send("BINOD")

    elif command.startswith("binod.news"):
        url = "https://newsapi.org/v2/top-headlines?sources=the-times-of-india&apiKey={}".format(os.environ['NEWS_API_KEY'])
        news = requests.get(url).text
        news_dict = json.loads(news)
        arts = news_dict['articles']

        a = 1
        string = "Today's news highlights are...\n"

        for article in arts:
            string += f"{article['title']}\n"

            if a == 2:
                break
            else:
                a += 1
                continue

        await message.channel.send(string)

    elif command.startswith("binod.weather"):
        query = command.replace("binod.weather", "")
        query = query.lstrip()
        query = query.capitalize()

        try:
            url = "http://api.weatherapi.com/v1/current.json?key={}&q={}".format(os.environ['WEATHER_API_KEY'], query)

            r = requests.get(url).text
            a = json.loads(r)
            current_weather = a.get('current')
            condition = current_weather['condition']
            text = condition['text']

            await message.channel.send(f"The current temperature in {query} is {current_weather['temp_c']}. The weather condition in {query} is {text.lower()}.")

        except Exception as e:
            print(e)
            await message.channel.send(f"Cannot search for a place named {query}!")

    elif command.startswith("binod.wiki"):
        try:
            await message.channel.send("Searching through the wikipedia for your query... Hang on...")
            query = command.replace('binod.wiki', "")
            query = query.lstrip()
            result = wikipedia.summary(query, sentences=2)
            material = f"According to wikipedia:\n{result}"
            await message.channel.send(material)

        except Exception as e:
            print(e)
            await message.channel.send("Some error occurred. Sorry for the inconvenience!")

    elif command.startswith("binod.corona"):
        country = command.replace('binod.corona', '').lstrip().rstrip()
        r = requests.get(
            f'https://api.covid19api.com/country/{country}/status/confirmed/live')

        try:
            parser = json.loads(r.text)
            l = parser[::-1]
            today = (l[0])
            material = (
                f"There are {today['Cases']} confirmed corona cases in {country.capitalize()}.")
            await message.channel.send(material)
        except Exception as e:
            print(e)
            await message.channel.send("Couldn't do that at the moment!")

    elif command.startswith('binod.joke'):
        types = ["knock-knock", 'jod', 'blonde', 'animal']
        url = f"https://api.jokes.one/jod?category={random.choice(types)}"
        r = requests.get(url).text
        parser = json.loads(r)
        joke = parser['contents']['jokes'][0]['joke']['text']
        await message.channel.send(f"Alright, here's a one:\n{joke}")

    elif command.startswith('binod.meme'):
        try:
            r = requests.get("https://imgflip.com/")
            soup = BeautifulSoup(r.content, "html5lib")

            images = soup.find_all("img")
            r = random.choice(images)
            link = (r["src"])
            ext = link.split(".")[-1]

            location = f"meme.{ext}"
            data = requests.get(("http:" + link)).content
            with open(location, "wb") as f:
                f.write(data)

            await message.channel.send(file=discord.File(location))

        except Exception as e:
            await message.channel.send("Some error occurred!")
            print(e)

    elif command.startswith("binod.roast"):
        target = command.replace("binod.roast", "").lstrip()
        if target.lower() == "shravan":
            await message.channel.send("How dare you mortal tryna roast my god!")
            await message.channel.send(insult(message.author.name))
        else:
            await message.channel.send(insult(target.capitalize()))

    elif command.startswith("binod.github"):
        query = command.replace("binod.github", "").lstrip().rstrip()
        if len(query.split("/")) == 1:
            await message.channel.send(github_user(query))
        elif len(query.split("/")) == 2:
            await message.channel.send(github_repo(query.split("/")[0], query.split("/")[1]))
        else:
            await message.channel.send("Invalid query!")


    elif command.startswith("/exec"):
        try:
            if message.author.name == "Shravan" and message.author.discriminator == "6942":
                query = command.replace("/exec", "").lstrip()
                output = subprocess.check_output(
                    query, shell=True).decode("utf-8")
                await message.channel.send(output)
            else:
                await message.channel.send("You're not allowed to do that!")

        except subprocess.CalledProcessError as e:
            await message.channel.send(e)


client.run(os.getenv("DISCORD_BOT_TOKEN"))
