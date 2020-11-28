from area import arearepo
from time import strftime
import re
import time
import pandas as pd
import statistics as st

class area_usecase:
    def getArea(self):
        areaUsecase = arearepo.area_repository()
        areaList = areaUsecase.getArea()
        exRate = areaUsecase.getExRate()

        if areaList == None or exRate == None:
            return False
        
        for area in areaList :
            try:
                price = float(area["price"])
                if price != None  : 
                    area["usd_price"] = str(price * exRate)         
            except:
                area["usd_price"] = "0"


        return areaList
    
    def getStatistics(self,info):
        areaUsecase = arearepo.area_repository()
        areaList = areaUsecase.getArea()      

        thetime = "2020-05-18"

        ts = time.strptime(thetime, '%Y-%m-%d')

        print(strftime("%U",ts))

        if areaList == None :
            return False

        statistic = []
        for area in areaList:
            if area["area_provinsi"].lower() == info["area_provinsi"].lower() :
                try:
                    price = float(area["price"])
                    tgl_parsed = time.strptime(area["tgl_parsed"][:10], '%Y-%m-%d')
                    tgl_match = time.strptime(str(info["week"][:10]), '%Y-%m-%d')
                    if (strftime("%U", tgl_parsed) == strftime("%U",tgl_match)):
                        statistic.append(price)
                except:
                    print("Some error on data")
        
        print(statistic)

        statistic_result = {
            "min" : min(statistic),
            "max" : max(statistic), 
            "average" : sum(statistic)/len(statistic),
            "median" : st.median(statistic)
        }

        return statistic_result