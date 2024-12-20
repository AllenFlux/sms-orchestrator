import requests

from sms_client import settings

def controller_sub_sms_report_list(page, sub_user_id):
    url = settings.url_prefix + "sms/sub/task/list"
    resp = requests.get(url, json={"pageNum": page, "sub_user_id": sub_user_id})
    if resp.status_code != 200:
        raise Exception(resp.status_code)
    return resp.json()

def controller_sub_sms_record_list(page, sub_user_id):
    url = settings.url_prefix + "sms/sub/task/record"
    resp = requests.get(url, json={"pageNum": page, "sub_user_id": sub_user_id})
    if resp.status_code != 200:
        raise Exception(resp.status_code)
    return resp.json()

def sub_get_conversation_record_list(sub_user_id, project_id, page):
    data = {
        "sub_user_id": sub_user_id,
        "project_id": project_id,
        "pageNum":page
    }
    url = settings.url_prefix + "sms/sub/conversation/record/list"
    resp = requests.get(url, json=data)
    if resp.status_code != 200:
        raise Exception(resp.status_code)
    return resp.json()