"""
{
"target_phone_number":["15122222222"],
"content":["Hi Hi Hi"]
}
"""
import json

from faker import Faker
fake = Faker(locale='en_US')

def generate_device_task_file():
    content_list = []
    phone_list = []
    for i in range(1000):
        content = fake.text(max_nb_chars=100)
        content_list.append(content)
        phone = fake.phone_number()
        phone_list.append(phone)
        print(content)
        print(phone)

    with open("device_task_file.json", "w") as f:
        json.dump({"target_phone_number": phone_list, "content": content_list}, f)


if __name__ == '__main__':
    generate_device_task_file()