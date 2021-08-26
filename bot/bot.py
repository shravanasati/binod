import json
import os
import random
import subprocess
import discord
import requests
import wikipedia
from src.insult import insult
from src.github import github_user, github_repo
from src.memes import get_meme
from src.shorten import shorten
from src.compile import CodeExecutor
from enchant.utils import levenshtein


client = discord.Client()
rce = CodeExecutor()


@client.event
async def on_ready():
    print('Bot has logged in as {0.user}'.format(client))





color = discord.Color.from_rgb(100, 230, 160)


@client.event
async def on_message(message):

    def cembed(title="", description="", thumbnail=client.user.avatar_url_as(format="png"), picture="", color=discord.Color.dark_theme()):
        embed = discord.Embed(title=title, description=description, color=color, thumbnail=thumbnail)
        return embed

    command = message.content.lower()

    if message.author == client.user:
        return

    help_message = '''
Hello, I am Binod. The following are the commands:
binod.help = Show this help message
binod.binod = BINOD 
binod.news = Top news headlines of the moment
binod.weather <place> = Real time weather of <place>  
binod.wiki <search> =Searches the wikipedia for <search>  
binod.joke = Get a random joke  
binod.corona = Get live status of confirmed corona cases in India  
binod.meme = Get a random meme  
binod.roast <object> = Roast <object>
binod.github <username>/<repo> = Get GitHub stats for a user or a repo
binod.shorten <url> = Shorten a URL
'''

    if command.startswith('binod.help'):
        await message.channel.send(embed=cembed(title="Help", description=help_message, color=color))

    elif command.startswith('binod.binod'):
        await message.channel.send("BINOD")

    elif command.startswith("binod.news"):
        url = "https://newsapi.org/v2/top-headlines?sources=the-times-of-india&apiKey={}".format(
            os.environ['NEWS_API_KEY'])
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

        await message.channel.send(embed=cembed(title="News", description=string, color=color))

    elif command.startswith("binod.weather"):
        query = command.replace("binod.weather", "")
        query = query.lstrip()
        query = query.capitalize()

        try:
            url = "http://api.weatherapi.com/v1/current.json?key={}&q={}".format(
                os.environ['WEATHER_API_KEY'], query)

            r = requests.get(url).text
            a = json.loads(r)
            current_weather = a.get('current')
            condition = current_weather['condition']
            text = condition['text']
            description = f"The current temperature in {query} is {current_weather['temp_c']}. The weather condition in {query} is {text.lower()}."

            await message.channel.send(embed = cembed(title=f"Weather of {query}", description=description, color=color))

        except Exception as e:
            print(e)
            await message.channel.send(embed=cembed(title="Weather", description=f"Cannot search for a place named {query}!", color=color))

    elif command.startswith("binod.wiki"):
        try:
            await message.channel.send(embed=cembed(description="Searching through the wikipedia for your query... Hang on...", color=color))
            query = command.replace('binod.wiki', "")
            query = query.lstrip()
            result = wikipedia.summary(query, sentences=2)
            material = f"According to wikipedia:\n{result}"
            await message.channel.send(embed=cembed(title="Wiki", description=material, color=color))

        except Exception as e:
            print(e)
            await message.channel.send(embed=cembed(title="Wiki", description="Some error occurred. Sorry for the inconvenience!", color=color))

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
            await message.channel.send(embed=cembed(title="Corona", description=material, color=color))
        except Exception as e:
            print(e)
            await message.channel.send("Couldn't do that at the moment!")

    elif command.startswith('binod.joke'):
        types = ["knock-knock", 'jod', 'blonde', 'animal']
        url = f"https://api.jokes.one/jod?category={random.choice(types)}"
        r = requests.get(url).text
        parser = json.loads(r)
        joke = parser['contents']['jokes'][0]['joke']['text']
        await message.channel.send(embed=cembed(title="Joke", description=f"Alright, here's a one:\n{joke}", color=color))

    elif command.startswith('binod.meme'):
        try:
            await message.channel.send(get_meme())

        except Exception as e:
            await message.channel.send("Some error occurred!")
            print(e)

    elif command.startswith("binod.roast"):
        target = command.replace("binod.roast", "").lstrip()
        if levenshtein(target.lower(), "shravan") <= 3 or "shravan" in target.lower():
            await message.channel.send("How dare you mortal tryna roast my senpai!")
            await message.channel.send(insult(message.author.name))
        else:
            await message.channel.send(insult(target.capitalize()))

    elif command.startswith("binod.github"):
        query = command.replace("binod.github", "").lstrip().rstrip()

        if len(query.split("/")) == 1:
            await message.channel.send(embed=cembed(title=f"GitHub Stats for {query}", description=github_user(query), color=color))

        elif len(query.split("/")) == 2:
            username = query.split("/")[0]
            repo = query.split("/")[1]

            await message.channel.send(embed=cembed(title=f"GitHub Stats for {query}", description=github_repo(username, repo), color=color))
        else:
            await message.channel.send(embed=cembed(title="GitHub Stats", description="Invalid query!", color=color))

    elif command.startswith("binod.shorten"):
        queries = command.replace(
            "binod.shorten", "").lstrip().rstrip().split(" ")
        if len(queries) == 1:
            e = cembed(title="Shorten", description=f"{shorten(queries[0])}", color=color)
            await message.channel.send(embed = e)
        elif len(queries) == 2:
            e = cembed(title="Shorten", description=f"{shorten(queries[0], queries[1])}", color=color)
            await message.channel.send(embed=e)
        elif len(queries) == 3:
            e = cembed(title="Shorten", description=f"{shorten(queries[0], queries[1], queries[2])}", color=color)
            await message.channel.send(embed=e)
        else:
            await message.channel.send("Invalid query! The command syntax is `binod.shorten <url_to_shorten> <alias_type> <alias>`.")

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

    elif command.startswith("binod.run"):
        split = command.replace("binod.run", "").lstrip().rstrip().split("\n")
        language = split[0].strip()

        actual_code = ""
        for i in split[1:]:
            if not i.startswith("```"):
                actual_code += i+"\n"

        resp = rce.execute_code(language, actual_code)
        result = ""
        if resp.exit_code == 0:
            result = f":white_check_mark: {message.author.mention}, the code executed successfully!\n"
        else:
            result = f":x: {message.author.mention}, the code didn't execute successfully!\n"
        result += f"```{resp.output}```"
        await message.channel.send(embed = cembed(title="Code Executed", description=result, color=color))

    elif command.startswith("binod.langs"):
        results = rce.runtimes
        resp_text = f"{message.author.mention}, here's the list of supported languages: \n"
        for i in results:
            lang = i["language"]
            version = i["version"]
            resp_text += f"{lang} v{version}\n"

        await message.channel.send(embed = cembed(title="Supported Languages", description=resp_text, color=color))

    elif command.startswith("binod.spam"):
        text = command.replace("binod.spam", "").lstrip().rstrip()
        if message.author.name == "Shravan" and message.author.discriminator == "6942":
            for i in range(10):
                print(i)
                await message.channel.send(text)
        else:
            await message.channel.send("You're not allowed to do that!")


client.run(os.getenv("DISCORD_BOT_TOKEN"))
