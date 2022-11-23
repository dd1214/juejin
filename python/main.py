'''
Description: 请填写文件简介
Version: 0.0
Autor: jlx
Date: 2022-08-25 16:00:24
LastEditors: jlx
LastEditTime: 2022-08-25 17:39:57
'''
import re
import json
import requests
import time
from bs4 import BeautifulSoup

# 推荐文章接口
url = "https://api.juejin.cn/recommend_api/v1/article/recommend_all_feed"
# 抓包 获取的参数
info = {"id_type": 2, "client_type": 2608, "sort_type": 200, "cursor": "0", "limit": 20}
headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36',
}
#  标记是否有更多数据
has_more = True

# 利用set 取一下重
dataSet = set()
# 存储数据的列表
dataList = []
# 需要的数据 数量
num = int(input("需要多少条文章数据："))

# 根据文章id获取文章的content
def get_content(aid):
    childUrl = f"https://juejin.cn/post/{aid}"
    childResp = requests.get(childUrl, headers=headers)
    childResp.encoding = "utf-8"
    content = childResp.text
    # print(content)

    res = re.findall('mark_content:"(.*?)",display_count:d', content)[0]
    res = re.sub(r'(\\u[\s\S]{4})', lambda x: x.group(1).encode("utf-8").decode("unicode-escape"), res)
    res = res.strip('"')
    return res


# 获取文章数据
def get_data():
    global has_more,num
    resp = requests.post(url, json=info, headers=headers)  # data自动使用json方式发送
    data = json.loads(resp.text)
    info["cursor"] = data["cursor"]
    has_more = data["has_more"]
    for i in data["data"]:
        item = i["item_info"]
        try:  # 去除广告推广内容的报错
            if item["article_id"] and len(dataSet) < num:
                if item["article_id"] in dataSet:
                    continue
                dataSet.add(item["article_id"])
                dic = {}
                content = get_content(item["article_id"])
                # id 封面 标题 简介 创建时间 修改事件 作者 头像 文章内容
                dic["article_id"] = item["article_id"]
                dic["cover_image"] = item["article_info"]["cover_image"]
                dic["title"] = item["article_info"]["title"]
                dic["brief"] = item["article_info"]["brief_content"]
                dic["ctime"] = item["article_info"]["ctime"]
                dic["mtime"] = item["article_info"]["mtime"]
                dic["user_name"] = item["author_user_info"]["user_name"]
                dic["avatar"] = item["author_user_info"]["avatar_large"]
                dic["content"] = content
                dataList.append(dic)

        except:
            continue
    print(f"当前{len(dataSet)}条数据已整理完毕！")
    time.sleep(4)



while has_more and len(dataSet) < num:
    get_data()

with open("article.json", "w", encoding="utf-8") as f:
    json.dump(dataList, f, indent=4, ensure_ascii=False)
f.close()

# 完成提示
print("over!!!")

# 最后结束提示
# 使用前需要进行 replace \\n  => \n 否则会出现乱码