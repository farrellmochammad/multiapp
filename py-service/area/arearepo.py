import requests
import requests_cache
import json
import pickle
from datetime import datetime, timedelta
import time


class area_repository:
    def getArea(self):
        requests_cache.install_cache(cache_name="pyservice-cache", backend='sqlite', expire_after=60)
        uri = "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"
        try:
            uResponse = requests.get(uri)
        except requests.ConnectionError:
            return None  
        Jresponse = uResponse.text
        data = json.loads(Jresponse)
        return data

    def getExRate(self):
        requests_cache.install_cache(cache_name="pyservice-cache", backend='sqlite', expire_after=60)
        uri = "https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=971e1c9087790b741a01"
        try:
            uResponse = requests.get(uri)
        except requests.ConnectionError:
            return None  
        Jresponse = uResponse.text
        data = json.loads(Jresponse)
        return data["IDR_USD"]
