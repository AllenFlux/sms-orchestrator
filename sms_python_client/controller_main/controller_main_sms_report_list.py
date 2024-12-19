import requests

from sms_client import settings

"""http://192.168.1.14:8822/api/v1/sms/task/list"""

def controller_main_sms_report_list(page):
    url = settings.url_prefix + "sms/task/list"
    resp = requests.get(url, json={"pageNum": page})
    if resp.status_code != 200:
        raise Exception(resp.status_code)
    return resp.json()