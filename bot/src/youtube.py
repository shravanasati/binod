import os
import requests
import json

def get_channel_info(channelID:str) -> str :
	API_KEY = os.environ.get('YOUTUBE_API_KEY')
	url = f"https://youtube.googleapis.com/youtube/v3/channels?part=snippet,contentDetails,statistics&id={channelID}&key={API_KEY}"

	response = requests.get(url)
	data = json.loads(response.text)
	if data.get("items") is None:
		return "Unable to find a youtube channel with {} ID.".format(channelID)

	items = data.get("items")[0]