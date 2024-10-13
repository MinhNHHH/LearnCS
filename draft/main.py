import requests
from bs4 import BeautifulSoup
import pandas as pd
import yaml
import threading
import re
# Title
# Company
# Location
# Link Company
# Description
# Email
# Salary

# Extended regular expression pattern to match various experience formats
pattern = r'(?:(hơn|trên|over|more than)\s*)?(\d+|một|hai|ba|bốn|năm|sáu|bảy|tám|chín|mười|\d+)\s*(năm|years?)\s*(kinh nghiệm|of experience)'

def format_url(url, param=None):
    if not param:
        return url
    url = url.format(**param)
    return url


def load_yml(file_path):
    config = {}
    with open(file_path, "r") as file:
        config = yaml.safe_load(file)

    return config


def regex(text, pattern):
    matches = re.findall(pattern, text, re.IGNORECASE)
    if matches:
        return matches.sort()
    return []


def fillter_condition(text, fillter):
   skills = fillter.get("skill", [])
   min_exp = fillter.get("min_exp", 0)

   # Find the number of years experience from text
   exp = regex(text, pattern)
   print(exp)
   if exp and len(exp) > 0 and exp[-1] >= min_exp:
       return True
   if any(skill.lower() in text.lower() for skill in skills):
       return True
    
   return False

def extract_requirements(url, configs):
    soup = parser_html(url)
    req = soup.find(configs.get("req_tag"), class_=configs.get("req_class"))
    if req:
        all_text = req.get_text(separator="\n", strip=True).split(configs.get("req_split"))
        if len(all_text) > 0:
           return all_text[-1]
    return ""

def statify_req(content, fillter):
    if content:
        statify_count = 0
        for text in content.split("\n"):
            if fillter_condition(text, fillter):
                statify_count += 1
        if statify_count >= len(fillter):
            return True
    return False

def parser_html(url):
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
        "Accept-Language": "en-US,en;q=0.9",
    }
    response = requests.get(url, headers=headers)
    return BeautifulSoup(response.text, "html.parser")


def get_href(soup):
    content = ""
    if soup:
        if soup.name == "a":
            content = soup.get("href")
        else:
            link = soup.find("a")
            if link:
                content = link.get("href")
    return content


def extract_data(url, configs):
    soup = parser_html(url)
    jobs = soup.find_all(
        configs.get("find_jobs_tag"), class_=configs.get("find_jobs_class")
    )
    data = []
    fillters = {
        "min_exp": 3,
        "skill": ["python", "go", "nodejs"]
    }
    if len(jobs) == 0:
        return data
    for job in jobs:
        title = job.find(configs.get("title_tag"), class_=configs.get("title_class"))
        company_name = job.find(
            configs.get("company_tag"), class_=configs.get("company_class")
        )
        location = job.find(
            configs.get("location_tag"), class_=configs.get("location_class")
        )
        salary = job.find(configs.get("salary_tag"), class_=configs.get("salary_class"))
        company_url = get_href(company_name)
        job_description_url = get_href(title)
        if statify_req(extract_requirements(job_description_url, configs), fillters):
            data.append(
                {
                    "Title": title.text.strip(),
                    "Company": company_name.text.strip(),
                    "Company_URL": company_url,
                    "Location": location.text.strip(),
                    "Job Description": job_description_url,
                    "Salary": salary.text.strip() if salary else "",
                }
            )
    return data


def process_platform(platform, config, data_lock, data_list):
    default_query = {'job': 'python'}
    url = config.get("url")
    local_data = []
    for i in range(10):
        if platform == "linkedin":
            default_query['range'] = i * 25
        else:
            default_query['page'] = i
        new_url = format_url(url, default_query)
        extract = extract_data(new_url, config)
        local_data.extend(extract)
        print(f"Done {new_url}")
    # Safely append to shared data_list using a lock to avoid race conditions
    with data_lock:
        data_list.extend(local_data)


def run_threads(configs):
    threads = []
    data = []
    data_lock = threading.Lock()  # Create a lock to synchronize data access
    # Create a thread for each platform
    for platform, config in configs.items():
        thread = threading.Thread(target=process_platform, args=(platform, config, data_lock, data))
        threads.append(thread)
        thread.start()  # Start the thread
    # Wait for all threads to complete
    for thread in threads:
        thread.join()

    return data


def main():
    configs = load_yml("config.yml")
    data = run_threads(configs)
    df = pd.DataFrame(data)
    df_no_duplicates = df.drop_duplicates()
    df_no_duplicates.to_excel("python.xlsx", index=False)

if __name__ == "__main__":
    main()
