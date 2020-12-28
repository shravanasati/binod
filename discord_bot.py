# TO DO:
# binod.commands done
# binod.news done
# binod.weather done
# binod.wikipedia done
# binod.corona done
# binod.ipl
# binod.joke done
# binod.meme done
# binod.suggest done

import discord, requests, json, wikipedia, os, random #, smtplib

client = discord.Client()

@client.event
async def on_ready():
    print('Bot has logged in as {0.user}'.format(client))

@client.event
async def on_message(message):
    command = message.content.lower()

    if message.author == client.user:
        return 

    help = "Hello I am a bot made by Shravan! You can do some really amazing things by following commands.\n\n binod.help = Bring this help message\n\n binod.binod = BINOD\n\n binod.news = Top news headlines of current time\n\n binod.weather <place> = Real time weather of <place> (Replace <place> with the place whose weather you want to know)\n\n binod.wikipedia <search> = Searches the wikipedia for <search> (Again replace the <search> with your query)\n\n binod.joke = Call a random joke \n\n binod.corona = Get live status of confirmed corona cases of India \n\n binod.meme = Get a random meme \n\n binod.suggest <suggestion> = Suggest a new feature to be added to the bot \n\n"

    if command.startswith('binod.help'):
        await message.channel.send(help)

    elif command.startswith('binod.binod'):
        await message.channel.send("BINOD")

    elif command.startswith("binod.news"):
        url = "https://newsapi.org/v2/top-headlines?sources=the-times-of-india&apiKey=fbddfcc479954110ba532f7b45dc41d2"
        news = requests.get(url).text
        news_dict = json.loads(news)
        arts = news_dict['articles']

        a = 1
        string = "Today's news highlights are...\n"

        for article in arts:
            string += f"{article['title']}\n"

            if a==2:
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
            url = f"http://api.weatherapi.com/v1/current.json?key=2f081a6878a747a5be135553200709&q={query}"

            r = requests.get(url).text
            a = json.loads(r)
            current_weather = a.get('current')
            condition = current_weather['condition']
            text = condition['text']

            await message.channel.send(f"The current tempearture in {query} is {current_weather['temp_c']}. The weather condition in {query} is {text.lower()}.")

        except Exception as e:
            print(e)
            await message.channel.send("Couldn't search that!")


    elif command.startswith("binod.wikipedia"):
        try:
            await message.channel.send("Searching through the wikipedia for your query... Hang on...")
            query = command.replace('binod.wikipedia', "")
            query = query.lstrip()
            result = wikipedia.summary(query, sentences = 2)
            material = f"According to wikipedia...\n{result}"
            await message.channel.send(material)

        except Exception as e:
            print(e)
            await message.channel.send("Some error occurred. Sorry for your inconvenience!")


    elif command.startswith("binod.corona"):
        r = requests.get(f'https://api.covid19api.com/country/india/status/confirmed/live')


        try:
            parser = json.loads(r.text)
            l = parser[::-1]
            today = (l[0])
            material = (f"There are {today['Cases']} confirmed corona cases in India.")
            await message.channel.send(material)
        except Exception as e:
            print(e)
            await message.channel.send("Couldn't do that at the moment!")

        
    elif command.startswith('binod.joke'):
        types = ["knock-knock", 'jod', 'blonde','animal']
        url = f"https://api.jokes.one/jod?category={random.choice(types)}"
        r = requests.get(url).text
        parser = json.loads(r)
        joke = parser['contents']['jokes'][0]['joke']['text']
        await message.channel.send(f"Alright, here's a one:\n{joke}")
        

    elif command.startswith('binod.meme'):
        meme_directory = r"C:\Users\Lenovo\Documents\Discord Bot Memes"
        contents = os.listdir(meme_directory)
        meme = random.choice(contents)
        await message.channel.send(file=discord.File(os.path.join(meme_directory, meme)))

    elif command.startswith("binod.suggest"):
        suggestion = command.replace("binod.suggest", "")
        author = message.author.mention
        author_name = message.author
        text = f"{author_name}: {suggestion}\n"
        with open(r"C:\Users\Lenovo\Documents\Python Codes\discord bot suggest.txt", 'a') as f:
            f.write(text)
        await message.channel.send(f"{author} Your suggestion has been recorded!")

    # elif command.startswith("binod.ask"):
    #     doubt = command.replace("binod.ask", "")
    #     author_mention = message.author.mention
    #     author = message.author
        
    #     subject = f"{author}'s doubt"
    #     body = f"Sir, \n{author} has following doubt regarding the python tutorials you're conducting. \n\n {author}: {doubt}"

    #     try:
    #         server = smtplib.SMTP('smtp.mail.yahoo.com', 587)
    #         server.ehlo()
    #         server.starttls()
    #         server.login('delunator.one@yahoo.com', "ytxyqkendtnngqib")
    #         server.sendmail('delunator.one@yahoo.com', 'delunator.one@yahoo.com', f"Subject: {subject}\n\n{body}")
    #         server.quit()
    #         await message.channel.send(f"{author_mention}, your doubt has been sent to Shravan.")

    #     except Exception as e:
    #         print(e)
    #         await message.channel.send("Couldn't connect at the moment! Try again later!")

client.run("NzYwNTE2MTc5MDQ2ODI1OTk0.X3NL4g.zRL8-FwDNSzL6YgXB2mbzHeWx6Y")