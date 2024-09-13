import requests
from bs4 import BeautifulSoup

# Title
# Company
# Location
# Link Company
# Description
# Email
# Salary
job = "golang"
location = "hanoi"
page = 1
url = f"https://www.topcv.vn/tim-viec-lam-{job}?sba={page}"
url1 = "https://vn.indeed.com/jobs?q=golang&l=&from=searchOnDesktopSerp&vjk=6e1058edc8e39cdb" 

def parser_html(url):
    response = requests.get(url)
    return BeautifulSoup(response.text, 'html.parser')

def extract_data(url, configs):
    soup = parser_html(url)
    jobs = soup.find_all(configs.get('find_jobs_tag'), class_=configs.get('find_jobs_class'))
    data = []
    print(soup)
    if len(jobs) == 0:
        return data
    # for job in jobs:
    #     title = job.find(configs.get('title_tag'), class_=configs.get('title_class'))
    #     company_name = job.find(configs.get('company_tag'), class_=configs.get('company_class'))
    #     location = job.find(configs.get('location_tag'), class_=configs.get('location_class'))
    #     salary = job.find(configs.get('salary_tag'), class_=configs.get('salary_class'))
    #     company_url = ""
    #     job_description_url = ""
    #     if title:
    #         url = title.find('a')
    #         if url:
    #             job_description_url = url.get('href')
    #     if company_name:
    #         company_url = company_name.get('href')
    #     data.append({
    #         "Title": title.text.strip(),
    #         "Company": company_name.text.strip(),
    #         "Company_URL": company_url,
    #         "Location": location.text.strip(),
    #         "Job Description": job_description_url,
    #         "Salary": salary.text.strip()
    #     })
    return data

def main():
    topcv = {
        "find_jobs_tag":"div",
        "find_jobs_class": "job-item-search-result",
        "title_tag":"h3",
        "title_class": "title",
        "company_tag": "a",
        "company_class": "company",
        "location_tag": "label",
        "location_class": "address",
        "salary_tag": "label",
        "salary_class": "title-salary",
    }

    linkedin = {
        "find_jobs_tag":"div",
        "find_jobs_class": "mosaic mosaic-provider-jobcards mosaic-provider-hydrated",
        "title_tag":"div",
        "title_class": "full-width artdeco-entity-lockup__title ember-view",
        "company_tag": "div",
        "company_class": "ember588",
        "location_tag": "div",
        "localtion_class": "d-flex align-items-center text-dark-grey imt-1",
        "salary_tag": "div",
        "salary_class": "salary",
    }
    data = extract_data(url1, linkedin)
    print(data)
    
if __name__=="__main__":
    main()
