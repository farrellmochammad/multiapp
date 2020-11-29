import requests
import requests_cache
import json
import pickle
from datetime import datetime, timedelta
import time
import os


class area_repository:
    def getArea(self):
        requests_cache.install_cache(cache_name="pyservice-cache", backend='sqlite', expire_after=60)
        uri = os.environ['STEIN_URL']
        try:
            uResponse = requests.get(uri)
        except requests.ConnectionError:
            return None  
        Jresponse = uResponse.text
        data = json.loads(Jresponse)
        return data

    def getExRate(self):
        requests_cache.install_cache(cache_name="pyservice-cache", backend='sqlite', expire_after=60)
        uri = os.environ['CONVERTER_URL']
        try:
            uResponse = requests.get(uri)
        except requests.ConnectionError:
            return None  
        Jresponse = uResponse.text
        data = json.loads(Jresponse)
        return data["IDR_USD"]
