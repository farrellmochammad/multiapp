import requests
import json
import pickle
from datetime import datetime, timedelta
import time


class area_repository:
    def __getArea__(self):
        uri = "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"
        try:
            uResponse = requests.get(uri)
        except requests.ConnectionError:
            return None  
        Jresponse = uResponse.text
        data = json.loads(Jresponse)
        return data

    def __getExRate__(self):
        uri = "https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=971e1c9087790b741a01"
        try:
            uResponse = requests.get(uri)
        except requests.ConnectionError:
            return None  
        Jresponse = uResponse.text
        data = json.loads(Jresponse)
        return data

    def getExRateCache(self):
        start_time = time.time()
        try :
            with open('./area/exrate.pkl', 'rb') as f:
                data = pickle.load(f)
            if (datetime.now() + timedelta(minutes=10)) < datetime.now():
                data = self.__getExRate__()
                with open('./area/exrate.pkl', 'wb') as f:
                    pickle.dump(data, f)  
        except Exception:
            data = self.__getExRate__()
            with open('./area/exrate.pkl', 'wb') as f:
                pickle.dump(data, f)
        return data["IDR_USD"]

    def getAreaCache(self):
        start_time = time.time()
        try :
            with open('./area/area.pkl', 'rb') as f:
                data = pickle.load(f)
            if (datetime.now() + timedelta(minutes=10)) < datetime.now():
                data = self.__getArea__()
                with open('./area/area.pkl', 'wb') as f:
                    pickle.dump(data, f)  
        except Exception:
            data = self.__getArea__()
            with open('./area/area.pkl', 'wb') as f:
                pickle.dump(data, f)
        
        return data
